package es

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
	"github.com/redis/go-redis/v9"
)

type InstanceMessageHandler func(header event.EventHeader, message InstanceMessage) bool
type TaskMessageHandler func(header event.EventHeader, message TaskMessage) bool

type redisEventHandler struct {
	rdb                    *redis.Client
	instanceMessageHandler InstanceMessageHandler
	taskMessageHandler     TaskMessageHandler

	maps map[string][]string
}

const (
	BidApplyProcessCode   = "PROC-958C3100-85BF-45D3-8583-6645DA922756" // 投标申请审批表单编码
	BidExpenseProcessCode = "PROC-D8453B77-B313-4BEB-BE42-C71EE81DA61A" // 投标项目付款表单编码
	TestProcessCode       = "PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97" // 数据拉取测试表单编码
)

func (reh *redisEventHandler) defaultHandler() EventHandler {
	return func(ctx context.Context, header event.EventHeader, rawData []byte, df payload.DataFrame) (event.EventProcessStatusType, error) {
		// 1. 幂等判断, 15s 内钉钉推送重复事件不处理
		if isEventProcessed(reh.rdb, header) {
			return event.EventProcessStatusKSuccess, nil
		}

		// 2. 判断是否支持的事件类型
		if !reh.isSupport(header, rawData) {
			return event.EventProcessStatusKSuccess, nil
		}

		// 3. 发送数据流
		values := streamValues(header, rawData)
		if err := reh.produceMessage(ctx, header, values); err != nil {
			return event.EventProcessStatusKLater, err
		}
		return event.EventProcessStatusKSuccess, nil
	}
}

func (reh *redisEventHandler) isSupport(header event.EventHeader, rawData []byte) bool {
	switch header.EventType {
	case "bpms_instance_change":
		instanceMessage := InstanceMessage{}
		_ = json.Unmarshal(rawData, &instanceMessage)
		if reh.instanceMessageHandler != nil {
			return reh.instanceMessageHandler(header, instanceMessage)
		}
		return true
	case "bpms_task_change":
		taskMessage := TaskMessage{}
		_ = json.Unmarshal(rawData, &taskMessage)
		if reh.taskMessageHandler != nil {
			return reh.taskMessageHandler(header, taskMessage)
		}
		return true
	}
	return false
}

// 1. 生产消息（XADD）：向 Stream 中添加消息
func (reh *redisEventHandler) produceMessage(ctx context.Context, header event.EventHeader, values []string) error {
	streamName := eventStreamKey(header)
	// XADD 语法：XADD 流名 * 字段1 值1 字段2 值2 ...
	// * 表示让 Redis 自动生成消息 ID（格式：时间戳-序列号，如 1735689600000-0）
	msgID, err := reh.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		ID:     "*", // 自动生成 ID
		Values: values,
	}).Result()
	if err != nil {
		// 如果生产消息失败，删除事件幂等 key
		reh.rdb.Del(ctx, eventStreamIdempotentKey(header))
	}
	fmt.Printf("生产消息成功 | 消息ID：%s\n", msgID)
	return nil
}

func RedisEventHandler(rdb *redis.Client) EventHandler {
	handler := &redisEventHandler{
		rdb: rdb,
		maps: map[string][]string{
			BidApplyProcessCode:   {"start", "finish", "terminate"},
			BidExpenseProcessCode: {"start", "finish", "terminate"},
			TestProcessCode:       {},
		}}
	return handler.defaultHandler()
}

// 幂等判断：检查事件是否已处理
func isEventProcessed(rdb *redis.Client, header event.EventHeader) bool {
	key := eventStreamIdempotentKey(header)
	// SETNX：仅当键不存在时设置，返回 true 表示未处理,
	ok, err := rdb.SetNX(context.TODO(), key, "1", 15*time.Second).Result()
	if err != nil {
		return false
	}
	// ok 为 false 表示键已存在，事件已处理
	return !ok
}

func eventStreamIdempotentKey(header event.EventHeader) string {
	return fmt.Sprintf("dingtalk:%s:event:%s", header.EventCorpId, header.EventId)
}

func eventStreamKey(header event.EventHeader) string {
	return fmt.Sprintf("dingtalk:%s:event:%s", header.EventCorpId, header.EventType)
}

func streamValues(header event.EventHeader, data []byte) []string {
	hb, _ := json.Marshal(header)
	values := []string{"header", string(hb), "data", string(data), "retry", "0"}
	return values
}
