package main

import (
	"cds/dingtalk/es"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
	"github.com/redis/go-redis/v9"
)

func main() {
	maps := map[string][]string{
		es.BidApplyProcessCode:   {"start", "finish", "terminate"},
		es.BidExpenseProcessCode: {"start", "finish", "terminate"},
		es.TestProcessCode:       {"start", "finish", "terminate", "delete"},
	}

	subscription := es.DingTalkEventSubscription(
		// 晶奇网关代理应用
		es.WithClient("dingygs46ockvmysbjlu", "t_o5NiKOA8Dy7wTtZ-wakzZ5-9Z-8u_JDH5hpXp7itk4cNouBfESswIp3-BuuffP"),
		// redis
		es.WithRedis(redis.Options{
			Addr:     "localhost:6379", // Redis 服务器地址
			Password: "redis!@@&",      // Redis 服务器密码
			DB:       2,                // Redis 数据库索引
		}),
		// 审批事件处理
		es.WithProcessInstanceEvent(func(header event.EventHeader, message es.InstanceMessage) bool {
			if val, ok := maps[message.ProcessCode]; ok {
				for _, v := range val {
					if strings.EqualFold(v, message.Type) {
						logger.GetLogger().Infof("process instance event: %v", message)
						return true
					}
				}
			}
			return false
		}),
		// 审批任务事件处理
		es.WithProcessTaskEvent(func(header event.EventHeader, message es.TaskMessage) bool {
			return false
		}),
		es.WithApproval(),
		es.WithEnt(),
	)

	err := subscription.Run(context.Background())
	if err != nil {
		panic(err)
	}

	// 注册路由，指定访问路径和处理函数
	//http.HandleFunc("/dict/", dictHandler)
	//
	//// 启动 HTTP 服务器，监听 8080 端口
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic("服务器启动失败: " + err.Error())
	//}

	//redis.Run()
}

// DictItem 定义字典项结构体，与返回数据结构对应
type DictItem struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Value int    `json:"value"`
	Sort  int    `json:"sort"`
}

// 字典数据处理函数
func dictHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 GET 方法
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("只支持 GET 方法"))
		return
	}

	// 解析路径中的 code（如从 /dict/RF 中提取 "RF"）
	path := r.URL.Path                         // 例如："/dict/RF"
	code := strings.TrimPrefix(path, "/dict/") // 去掉前缀 "/dict/"，得到 "RF"

	// 校验 code 是否为空（如访问 /dict/ 时会触发）
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("请指定字典编码，例如 /dict/RF"))
		return
	}

	// 设置响应头为 JSON 格式
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// 定义要返回的字典数据
	dictData := []DictItem{
		{
			Code:  "RF",
			Name:  "报名费",
			Value: 1,
			Sort:  1,
		},
		{
			Code:  "DF",
			Name:  "标书工本费",
			Value: 2,
			Sort:  2,
		},
	}

	// 将数据序列化为 JSON 并写入响应
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(dictData); err != nil {
		// 处理序列化错误
		http.Error(w, "数据序列化失败", http.StatusInternalServerError)
		return
	}
}
