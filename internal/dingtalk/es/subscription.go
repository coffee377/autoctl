package es

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/card"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/plugin"
	"github.com/redis/go-redis/v9"
)

type Subscription interface {
	Run(ctx context.Context) error
}

func DingTalkEventSubscription(options ...SubscriptionOption) Subscription {
	s := &subscription{}
	for _, opt := range options {
		opt(s)
	}
	if s.rdb == nil {
		s.rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	}
	if s.eventFrameHandler == nil {
		s.eventFrameHandler = NewRedisEventFrameHandler(s.rdb)
	}
	return s
}

type subscription struct {
	clientId     string
	clientSecret string

	rdb               *redis.Client
	eventFrameHandler *RedisEventFrameHandler
}

const (
	corpId             = "dingd8b32bfb2b9da7b2"
	bpmsInstanceChange = "bpms_instance_change"
	bpmsTaskChange     = "bpms_task_change"
)

func (s *subscription) Run(ctx context.Context) error {

	logger.SetLogger(logger.NewStdTestLogger())

	cli := client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(s.clientId, s.clientSecret)))

	// 注册事件类型的处理函数
	if s.eventFrameHandler != nil {
		cli.RegisterAllEventRouter(s.eventFrameHandler.OnEventReceived)
	}

	// 注册机器人 callback 类型的处理函数
	// cli.RegisterChatBotCallbackRouter(OnChatBotMessageReceived)

	// 注册 AI 插件的处理函数
	// cli.RegisterPluginCallbackRouter(OnPluginMessageReceived)

	// 注册互动卡片类型的处理函数
	// cli.RegisterCardCallbackRouter(OnCardCallbackReceived)

	err := cli.Start(ctx)
	if err != nil {
		return err
	}
	defer cli.Close()

	bpmsInstanceChangeStream := fmt.Sprintf("dingtalk:%s:event:%s", corpId, bpmsInstanceChange)
	bpmsTaskChangeStream := fmt.Sprintf("dingtalk:%s:event:%s", corpId, bpmsTaskChange)
	// 创建主消费组
	s.createConsumerGroup(ctx, bpmsInstanceChangeStream, 2)
	s.createConsumerGroup(ctx, bpmsTaskChangeStream, 2)

	// 创建死信队列消费组
	//s.createConsumerGroup(ctx, fmt.Sprintf("dingtalk:dlq:%s:%s", corpId, bpmsInstanceChange), bpmsInstanceChange)
	//s.createConsumerGroup(ctx, fmt.Sprintf("dingtalk:dlq:%s:%s", corpId, bpmsTaskChange), bpmsTaskChange)

	go s.cleanupStreamMessages(context.Background(), cleanupInterval, bpmsInstanceChangeStream, bpmsTaskChangeStream)

	go s.consumeMessage(context.Background(), bpmsInstanceChangeStream, "group_1", "consumer_01")
	go s.consumeMessage(context.Background(), bpmsTaskChangeStream, "group_2", "consumer_02")

	select {}
}

// 2. 创建消费组（XGROUP CREATE）：首次消费前必须创建消费组
func (s *subscription) createConsumerGroup(ctx context.Context, stream string, groupNum int8) {
	// XGROUP CREATE 语法：XGROUP CREATE 流名 消费组名 消息ID [MKSTREAM]
	// 消息ID：$ 表示从最新消息开始消费；0-0 表示从最早消息开始消费
	// MKSTREAM：若流不存在则自动创建流（避免报错）
	for i := 0; i < int(groupNum); i++ {
		group := fmt.Sprintf("group_%d", i+1)
		s.rdb.XGroupCreateMkStream(ctx, stream, group, "0-0")
	}

	//if err != nil {
	//	if err.Error() != "BUSYGROUP Consumer Group name already exists" {
	//		fmt.Printf("创建消费组失败（流：%s，组：%s）：%v\n", stream, group, err)
	//	} else {
	//		fmt.Printf("消费组已存在（流：%s，组：%s）\n", stream, group)
	//	}
	//} else {
	//	fmt.Printf("消费组创建成功（流：%s，组：%s）\n", stream, group)
	//}
}

// 3. 消费组消费（XREADGROUP）：多个消费者分摊消费，需 ACK 确认
func (s *subscription) consumeMessage(ctx context.Context, streamName, groupName, consumerName string) {
	fmt.Printf("\n消费者 %s 开始监听消息...\n", consumerName)
	for {
		// 1. 消费未处理的消息（">"）
		// > 表示消费组中未被消费的消息；若要处理 Pending 消息，可替换为具体消息ID或0
		streams, err := s.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			//Block:    2 * time.Second, // 阻塞2秒（避免永久阻塞，方便后续扫描Pending）
			Block:   0, // 阻塞直到有新消息
			Count:   1,
			Streams: []string{streamName, ">"},
		}).Result()

		if err != nil && !errors.Is(err, redis.Nil) {
			fmt.Printf("消费者 %s 读取消息失败：%v\n", consumerName, err)
			time.Sleep(1 * time.Second)
			continue
		}

		// 2. 处理读取到的消息
		hasNewMsg := false
		for _, stream := range streams {
			for _, msg := range stream.Messages {
				hasNewMsg = true
				// 2.1 处理消息（核心：重试逻辑 + 死信队列）
				s.processSingleMsg(ctx, streamName, groupName, consumerName, msg)
			}
		}

		// 3. 若无新消息，扫描Pending消息重试（每轮阻塞后执行一次）
		if !hasNewMsg {
			s.scanAndRetryPendingMessage(ctx, consumerName, streamName, groupName)
		}
	}
}

const (
	maxRetryCount   = 3               // 最大重试次数
	cleanupInterval = 5 * time.Minute // 消息清理间隔（5分钟）
	retainHours     = 24              // 保留消息时长（超过24小时则清理）
)

// 处理单条消息（核心：重试逻辑 + 死信队列）
func (s *subscription) processSingleMsg(ctx context.Context, streamName, groupName, consumerName string, msg redis.XMessage) {
	// 解析消息内容
	var (
		header   event.EventHeader
		data     interface{}
		retryCnt int
	)
	e1 := json.NewDecoder(strings.NewReader(msg.Values["header"].(string))).Decode(&header)
	e2 := json.NewDecoder(strings.NewReader(msg.Values["data"].(string))).Decode(&data)
	if e1 != nil || e2 != nil {
		return
	}
	i := msg.Values["retry"]
	switch v := i.(type) {
	case string:
		retryCnt, _ = strconv.Atoi(v) // 读取重试次数
	default:
		retryCnt = 0
	}
	msgID := msg.ID

	fmt.Printf("消费者 %s 接收消息 | 消息ID：%s | 重试次数：%d\n", consumerName, msgID, retryCnt)
	//fmt.Printf("Header => %v \n", msg.Values["header"])
	//fmt.Printf("Data => %v \n", msg.Values["data"])

	bid := false
	switch d := data.(type) {
	case InstanceMessage:
		bid = d.ProcessCode == BidApplyProcessCode || d.ProcessCode == BidExpenseProcessCode
	case TaskMessage:
		bid = d.ProcessCode == BidApplyProcessCode || d.ProcessCode == BidExpenseProcessCode
	}

	// 模拟不处理投标数据（故意让部分消息失败，测试重试）
	if bid {
		fmt.Printf("消费者 %s 处理消息失败（ID：%s），进入重试\n", consumerName, msgID)
		// 重试次数+1，更新消息元数据（通过 XADD 重新加入Pending列表，原消息需删除）
		retryCnt++
		if retryCnt > maxRetryCount {
			// 超过最大重试次数，移至死信队列
			//		moveToDeadLetter(msg)
			//fmt.Printf("消息ID：%s 重试超过上限（%d次），移至死信队列\n", msgID, maxRetryCount)
			// ACK 原消息（从主流Pending列表移除）
			//s.rdb.XAck(ctx, streamName, groupName, msgID)
		} else {
			// 未超过重试次数，更新重试次数后重新加入流（继续Pending）
			s.rdb.XAdd(ctx, &redis.XAddArgs{
				Stream: streamName,
				ID:     "*", // 重新生成ID（避免冲突）
				Values: map[string]interface{}{
					"header":    msg.Values["header"],
					"data":      msg.Values["data"],
					"retry_cnt": retryCnt,
				},
			})
			// ACK 原消息（删除旧的失败消息）
			s.rdb.XAck(ctx, streamName, groupName, msgID)
		}
		return
	}

	// 处理成功，ACK 确认
	err := s.rdb.XAck(ctx, streamName, groupName, msgID).Err()
	if err != nil {
		fmt.Printf("消费者 %s ACK 消息失败（ID：%s）：%v\n", consumerName, msgID, err)
	} else {
		fmt.Printf("消费者 %s 处理并 ACK 消息成功（ID：%s）\n", consumerName, msgID)
	}
}

// 扫描并重试 Pending 消息（防止消息遗漏）
func (s *subscription) scanAndRetryPendingMessage(ctx context.Context, consumerName, streamName, groupName string) {
	// XPENDING 获取 Pending 消息列表（只取前5条，避免一次性处理过多）
	pendingMsg, err := s.rdb.XPending(ctx, streamName, groupName).Result()

	if err != nil || pendingMsg.Count == 0 {
		return
	}

	fmt.Printf("\n消费者 %s 扫描到 %d 条 Pending 消息，开始重试...\n", consumerName, pendingMsg.Count)

	//for k, _ := range pendingMsg.Consumers {
	//	// XCLAIM：将Pending消息转移到当前消费者，重新处理
	//	claimedMsg, err := rdb.XClaim(ctx, &redis.XClaimArgs{
	//		Stream:   streamName,
	//		Group:    groupName,
	//		Consumer: consumerName,
	//		MinIdle:  1 * time.Minute, // 消息闲置1分钟后再重试（避免并发冲突）
	//		Messages: []string{k},
	//	}).Result()
	//
	//	if err != nil || len(claimedMsg) == 0 {
	//		continue
	//	}
	//
	//	// 处理转移后的消息
	//	for _, msg := range claimedMsg {
	//		processSingleMsg(consumerName, &msg)
	//	}
	//}
}

// 移至死信队列
func (s *subscription) moveToDeadLetter(ctx context.Context, msg *redis.XMessage) {
	// 将消息添加到死信队列
	//_, err := s.rdb.XAdd(ctx, &redis.XAddArgs{
	//	Stream: deadLetterStream,
	//	ID:     "*",
	//	Values: msg.Values,
	//}).Result()
	//
	//if err != nil {
	//	fmt.Printf("消息ID：%s 移至死信队列失败：%v\n", msg.ID, err)
	//}
}

// 流消息清理策略（定时执行）
func (s *subscription) cleanupStreamMessages(ctx context.Context, cleanupInterval time.Duration, streamName string, streamNames ...string) {
	ticker := time.NewTicker(cleanupInterval)
	defer ticker.Stop()

	cleanup := func(ctx context.Context, streamName string) {
		if streamName == "" {
			return
		}
		result, _ := s.rdb.XInfoGroups(ctx, streamName).Result()
		lastDeliveredIDs := make([]string, 0)
		for _, group := range result {
			if group.Consumers > 0 {
				lastDeliveredIDs = append(lastDeliveredIDs, group.LastDeliveredID)
			}
		}
		if minID, ok := MinStreamID(lastDeliveredIDs...); ok {
			fmt.Printf("清理流：%s，最小ID：%s\n", streamName, minID)
			s.rdb.XTrimMinIDApprox(ctx, streamName, minID, 100)
		}
	}

	fmt.Printf("\n消息清理任务启动，每 %v 执行一次\n", cleanupInterval)

	for range ticker.C {
		fmt.Println("\n=== 开始执行消息清理 ===")
		// 2. 按时间清理：保留最近 retainHours 小时的消息
		// 计算截止时间戳（当前时间 - retainHours 小时），单位：毫秒
		//cutoffTimestamp := time.Now().Add(-5 * time.Minute).UnixMilli()
		//minID := fmt.Sprintf("%d-0", cutoffTimestamp) // Stream ID 格式：时间戳-序列号
		cleanup(ctx, streamName)
		for _, name := range streamNames {
			cleanup(ctx, name)
		}
		fmt.Println("=== 消息清理执行结束 ===")
	}
}

// OnChatBotMessageReceived 简单的应答机器人实现
func OnChatBotMessageReceived(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	replyMsg := []byte(fmt.Sprintf("msg received: [%s]", data.Text.Content))

	chatbotReplier := chatbot.NewChatbotReplier()
	if err := chatbotReplier.SimpleReplyText(ctx, data.SessionWebhook, replyMsg); err != nil {
		return nil, err
	}
	//if err := chatbotReplier.SimpleReplyMarkdown(ctx, data.SessionWebhook, []byte("Markdown消息"), replyMsg); err != nil {
	//	return nil, err
	//}

	//SimpleReplyCard(ctx, data.SessionWebhook, chatbotReplier, []byte("Markdown消息"), replyMsg)

	return replyMsg, nil
}

// OnPluginMessageReceived 简单的插件处理实现
func OnPluginMessageReceived(ctx context.Context, request *plugin.GraphRequest) (*plugin.GraphResponse, error) {
	response := &plugin.GraphResponse{
		Body: `{"text": "hello world", "content": [{"title": "1", "description": "2", "url":"https://www.zhihu.com/question/626551401"},{"title": "2", "description": "2", "url":"https://www.zhihu.com/question/626551401"}]}`,
	}
	return response, nil
}

// OnCardCallbackReceived 卡片回调处理
func OnCardCallbackReceived(ctx context.Context, request *card.CardRequest) (*card.CardResponse, error) {
	logger.GetLogger().Infof("receive card data: %v", request)
	action := request.CardActionData.CardPrivateData.Params["action"]
	logger.GetLogger().Infof("action: %s", action)

	cardParamMap := map[string]string{}
	switch action {
	case "build":
		cardParamMap["buildable"] = "false"
		cardParamMap["deployable"] = "true"
		cardParamMap["reversible"] = "false"
	case "deploy":
		cardParamMap["buildable"] = "false"
		cardParamMap["deployable"] = "false"
		cardParamMap["reversible"] = "true"
	case "rollback":
		cardParamMap["buildable"] = "false"
		cardParamMap["deployable"] = "false"
		cardParamMap["reversible"] = "false"
	}

	response := &card.CardResponse{
		CardUpdateOptions: &card.CardUpdateOptions{
			UpdateCardDataByKey:    false,
			UpdatePrivateDataByKey: true,
		},
		CardData: &card.CardDataDto{
			CardParamMap: map[string]string{},
		},
		UserPrivateData: &card.CardDataDto{
			CardParamMap: cardParamMap,
		},
	}
	return response, nil
}
