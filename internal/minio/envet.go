package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/notification"
	"log"
	"net/http"
)

type MinIOEvent struct {
	EventVersion string        `json:"EventVersion"`
	Events       []EventDetail `json:"Events"`
}

type EventDetail struct {
	EventName string   `json:"EventName"` // 事件类型（如"s3:ObjectCreated:Put"）
	Key       string   `json:"Key"`       // 文件路径
	Records   []Record `json:"Records"`
}

type Record struct {
	S3        S3Detail `json:"s3"`
	EventTime string   `json:"eventTime"` // 事件时间
}

type S3Detail struct {
	Bucket BucketDetail `json:"bucket"`
	Object ObjectDetail `json:"object"`
}

type BucketDetail struct {
	Name string `json:"name"` // 桶名
}

type ObjectDetail struct {
	Key  string `json:"key"`  // 文件路径
	Size int64  `json:"size"` // 文件大小（字节）
}

// 处理MinIO事件的Webhook端点
func handleMinIOEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体（事件JSON）
	var event EventDetail
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		fmt.Printf("解析事件失败: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 提取事件关键信息（示例：处理第一个事件）
	if len(event.Records) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	fmt.Printf("收到事件: %s\n", event.EventName)

	fmt.Printf("文件路径: %s\n", event.Key)
	fmt.Printf("所在桶: %s\n", event.Records[0].S3.Bucket.Name)
	fmt.Printf("文件大小: %d字节\n", event.Records[0].S3.Object.Size)
	fmt.Printf("事件时间: %s\n", event.Records[0].EventTime)

	// 执行后续业务逻辑（如文件解析、通知等）
	// processFile(detail.Records[0].S3.Bucket.Name, detail.Key)

	// 必须返回200 OK，否则MinIO会重试
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func main() {
	// 注册Webhook端点
	http.HandleFunc("/minio-events", handleMinIOEvent)

	// 启动服务（端口8080，需与配置的Webhook端点一致）
	fmt.Println("Webhook服务启动，监听端口9098...")
	if err := http.ListenAndServe(":9098", nil); err != nil {
		fmt.Printf("服务启动失败: %v\n", err)
	}
}

func main2() {
	// 初始化MinIO客户端
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("console", "console@1227", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("初始化MinIO失败: %v", err)
	}

	// 配置事件通知
	if err := setupMinIOEventNotification(minioClient, "download", "http://localhost:8080/minio-events"); err != nil {
		log.Fatalf("配置事件通知失败: %v", err)
	}

	// 注册Webhook端点
	http.HandleFunc("/minio-events", func(w http.ResponseWriter, r *http.Request) {
		handleMinIOEvent2(minioClient, w, r)
	})

	// 启动服务
	log.Println("Webhook服务启动，监听端口8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 处理MinIO事件的Webhook端点
func handleMinIOEvent2(minioClient *minio.Client, w http.ResponseWriter, r *http.Request) {
	// 解析事件JSON
	var event MinIOEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "解析事件失败", http.StatusBadRequest)
		return
	}

	// 提取文件信息
	if len(event.Events) == 0 {
		http.Error(w, "无事件数据", http.StatusOK)
		return
	}
	//detail := event.Events[0]
	//bucketName := detail.Records[0].S3.Bucket.Name
	//objectKey := detail.Key

	//// 从MinIO下载文件
	//fileData, err := downloadMinIOFile(minioClient, bucketName, objectKey)
	//if err != nil {
	//	http.Error(w, "下载文件失败", http.StatusInternalServerError)
	//	return
	//}
	//
	//// 通过钉钉API发送文件
	//if err := sendFileToDingTalk(fileData, objectKey); err != nil {
	//	http.Error(w, "钉钉文件发送失败", http.StatusInternalServerError)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

// 配置MinIO事件通知
func setupMinIOEventNotification(minioClient *minio.Client, bucketName, webhookURL string) error {
	ctx := context.Background()

	//// 定义监听事件类型（如文件上传）
	//events := event.NewEvents([]event.Name{event.ObjectCreatedPut})

	//// 配置Webhook目标（ARN格式固定为"arn:minio:sqs::1:webhook"）
	//queueConfig := NewQueueConfig(
	//	event.ARN{Value: "arn:minio:sqs::1:webhook"},
	//	events,
	//).WithEndpoint(webhookURL) // Webhook端点
	//
	//// 应用配置到桶
	//notificationConfig := event.NewNotificationConfig([]event.QueueConfig{queueConfig})
	//notificationConfig := notifications
	queueArn := notification.NewArn("minio", "sqs", "", "primary", "webhook")
	queueArn, _ = notification.NewArnFromString("arn:minio:sqs::1:webhook")
	queueConfig := notification.NewConfig(queueArn)
	queueConfig.AddEvents(notification.ObjectCreatedPut, notification.ObjectRemovedDelete)
	//queueConfig.AddFilterPrefix("photos/")
	queueConfig.AddFilterSuffix(".xlsx")

	notificationConfig := notification.Configuration{}
	notificationConfig.AddQueue(queueConfig)
	return minioClient.SetBucketNotification(ctx, bucketName, notificationConfig)
}
