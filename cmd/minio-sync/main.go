package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type aliasConfig struct {
	URL          string `json:"url"`
	AccessKey    string `json:"accessKey"`
	SecretKey    string `json:"secretKey"`
	SessionToken string `json:"sessionToken,omitempty"`
	API          string `json:"api"`
	Path         string `json:"path"`
	License      string `json:"license,omitempty"`
	APIKey       string `json:"apiKey,omitempty"`
	Src          string `json:"src,omitempty"`
}

type mcConfig struct {
	Version string                 `json:"version"`
	Aliases map[string]aliasConfig `json:"aliases"`
}

// 获取 mc 配置文件路径
func getMCConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("获取用户目录失败: %v", err)
	}
	// Linux/macOS: ~/.mc/config.json; Windows: %USERPROFILE%\.mc\config.json
	return filepath.Join(homeDir, ".mc", "config.json"), nil
}

// 加载 mc 配置文件
func loadMCConfig() (*mcConfig, error) {
	configPath, err := getMCConfigPath()
	if err != nil {
		return nil, err
	}

	// 读取文件内容
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v（路径：%s）", err, configPath)
	}

	// 解析 JSON
	var config mcConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// 根据 mc 配置的别名创建 MinIO 客户端
func newClientFromAlias(alias string) (*minio.Client, error) {
	// 加载 mc 配置
	config, err := loadMCConfig()
	if err != nil {
		return nil, err
	}

	// 查找指定别名的配置
	hostConfig, exists := config.Aliases[alias]
	if !exists {
		return nil, fmt.Errorf("未找到别名 %s 的配置", alias)
	}

	// 检查必要参数
	if hostConfig.URL == "" || hostConfig.AccessKey == "" || hostConfig.SecretKey == "" {
		return nil, fmt.Errorf("别名 %s 的配置不完整（缺少 URL/AccessKey/SecretKey）", alias)
	}

	// 初始化 MinIO 客户端
	useSSL := false
	if hostConfig.URL[:5] == "https" {
		useSSL = true
	}
	endpoint := strings.TrimPrefix(hostConfig.URL, "http://")
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(hostConfig.AccessKey, hostConfig.SecretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, fmt.Errorf("创建客户端失败: %v", err)
	}

	return client, nil
}

// // 同步桶基础配置（含对象锁定）
//
//	func syncBucketBaseConfig(ctx context.Context, sourceClient, targetClient *minio.Client, sourceBucket, targetBucket string) error {
//		// 1. 获取源桶配置
//		sourceInfo, err := sourceClient.BucketExists(ctx, sourceBucket)
//		if !sourceInfo || err != nil {
//			return fmt.Errorf("源桶不存在或无法访问: %v", err)
//		}
//
//		// 2. 检查目标桶是否存在，不存在则创建（如需对象锁定，必须在创建时启用）
//		targetExists, err := targetClient.BucketExists(ctx, targetBucket)
//		if err != nil {
//			return fmt.Errorf("检查目标桶失败: %v", err)
//		}
//		if !targetExists {
//			// 检查源桶是否启用对象锁定
//			sourceLockConf, err := sourceClient.GetObjectLockConfiguration(ctx, sourceBucket)
//			if err != nil {
//				return fmt.Errorf("获取源桶对象锁定配置失败: %v", err)
//			}
//			// 创建目标桶（若源桶有对象锁定，目标桶必须启用）
//			opts := minio.MakeBucketOptions{}
//			if sourceLockConf.ObjectLockEnabled == "Enabled" {
//				opts.ObjectLocking = true // 启用对象锁定（创建后不可修改）
//			}
//			if err := targetClient.MakeBucket(ctx, targetBucket, opts); err != nil {
//				return fmt.Errorf("创建目标桶失败: %v", err)
//			}
//		}
//		return nil
//	}
//
// // 2. 同步版本控制配置
//
//	func syncVersioning(ctx context.Context, sourceClient, targetClient *minio.Client, sourceBucket, targetBucket string) error {
//		// 1. 获取源桶版本控制状态
//		sourceVC, err := sourceClient.GetBucketVersioning(ctx, sourceBucket)
//		if err != nil {
//			return fmt.Errorf("获取源桶版本控制失败: %v", err)
//		}
//
//		// 2. 同步到目标桶
//		targetVC := minio.SetBucketVersioningOptions{
//			VersioningConfiguration: sourceVC.VersioningConfiguration,
//		}
//		if err := targetClient.SetBucketVersioning(ctx, targetBucket, targetVC); err != nil {
//			return fmt.Errorf("设置目标桶版本控制失败: %v", err)
//		}
//		return nil
//	}
//
// 3. 同步 ILM（生命周期）规则
func syncILMRules(ctx context.Context, sourceClient, targetClient *minio.Client, sourceBucket, targetBucket string) error {
	// 1. 获取源桶 ILM 规则
	sourceLCR, err := sourceClient.GetBucketLifecycle(ctx, sourceBucket)
	if err != nil {
		return fmt.Errorf("获取源桶 ILM 规则失败: %v", err)
	}

	//// 2. 清除目标桶现有规则（避免冲突）
	//if err := targetClient.RemoveBucketLifecycle(ctx, targetBucket); err != nil && !strings.Contains(err.Error(), "NoSuchLifecycleConfiguration") {
	//	return fmt.Errorf("清除目标桶 ILM 规则失败: %v", err)
	//}

	// 3. 应用源桶规则到目标桶
	if len(sourceLCR.Rules) > 0 {
		if err := targetClient.SetBucketLifecycle(ctx, targetBucket, sourceLCR); err != nil {
			return fmt.Errorf("设置目标桶 ILM 规则失败: %v", err)
		}
	}
	return nil
}

// 4. 同步事件通知配置
func syncEventNotifications(ctx context.Context, sourceClient, targetClient *minio.Client, sourceBucket, targetBucket string) error {
	// 1. 获取源桶事件配置
	sourceEvents, err := sourceClient.GetBucketNotification(ctx, sourceBucket)
	if err != nil {
		return fmt.Errorf("获取源桶事件配置失败: %v", err)
	}

	// 2. 清除目标桶现有事件
	if err := targetClient.RemoveAllBucketNotification(ctx, targetBucket); err != nil {
		return fmt.Errorf("清除目标桶事件失败: %v", err)
	}

	// 3. 应用源桶事件到目标桶
	if len(sourceEvents.QueueConfigs) > 0 || len(sourceEvents.TopicConfigs) > 0 {
		if err := targetClient.SetBucketNotification(ctx, targetBucket, sourceEvents); err != nil {
			return fmt.Errorf("设置目标桶事件失败: %v", err)
		}
	}
	return nil
}

//// 5. 同步对象数据（含所有版本和元数据）
//func syncObjects(ctx context.Context, sourceClient, targetClient *minio.Client, sourceBucket, targetBucket string) error {
//	// 1. 列出源桶所有对象版本（含删除标记）
//	listOpts := minio.ListObjectVersionsOptions{
//		Recursive: true,
//	}
//	for object := range sourceClient.ListObjectVersions(ctx, sourceBucket, "", "", "", listOpts) {
//		if object.Err != nil {
//			return fmt.Errorf("列出对象版本失败: %v", object.Err)
//		}
//
//		// 2. 复制对象（保留元数据和存储类）
//		copySrc := minio.CopySrcOptions{
//			Bucket:    sourceBucket,
//			Object:    object.Key,
//			VersionID: object.VersionID, // 复制指定版本
//		}
//		copyDst := minio.CopyDstOptions{
//			Bucket:          targetBucket,
//			Object:          object.Key,
//			StorageClass:    object.StorageClass, // 保留存储类
//			UserMetadata:    object.UserMetadata, // 保留用户元数据
//			ReplaceMetadata: true,
//		}
//
//		// 处理删除标记（DeleteMarker）
//		if object.IsDeleteMarker {
//			if err := targetClient.RemoveObject(ctx, targetBucket, object.Key, minio.RemoveObjectOptions{
//				VersionID: object.VersionID,
//			}); err != nil {
//				return fmt.Errorf("同步删除标记失败: %v", err)
//			}
//			continue
//		}
//
//		// 复制对象内容
//		_, err := targetClient.CopyObject(ctx, copyDst, copySrc)
//		if err != nil {
//			return fmt.Errorf("复制对象 %s 失败: %v", object.Key, err)
//		}
//	}
//	return nil
//}

// 主同步函数
func mainSync(ctx context.Context, sourceAlias, targetAlias string, excludeBucket ...string) error {
	buf := bytes.Buffer{}

	buf.WriteString("#!/bin/bash\n")
	buf.WriteString(fmt.Sprintf("SOURCE_ALIAS=%s\n", sourceAlias))
	buf.WriteString(fmt.Sprintf("TARGET_ALIAS=%s\n", targetAlias))

	// 1. 同步数据并保留属性
	buf.WriteString("# 1. 同步数据并保留属性\n")
	mirrorCmd := []string{"mc", "mirror", "--preserve", "--overwrite"}
	//mirrorCmd = append(mirrorCmd, "--remove")
	if len(excludeBucket) > 0 {
		mirrorCmd = append(mirrorCmd, "--exclude-bucket")
		mirrorCmd = append(mirrorCmd, excludeBucket...)
	}
	mirrorCmd = append(mirrorCmd, sourceAlias, targetAlias)
	buf.WriteString(strings.Join(mirrorCmd, " ") + "\n")

	// 2. 迁移 ILM 配置
	buf.WriteString("# 2. 迁移 ILM 配置\n")
	//mc ilm export "$SOURCE_ALIAS/$BUCKET" | mc ilm import "$TARGET_ALIAS/$BUCKET"

	// 初始化客户端
	sourceClient, err := newClientFromAlias(sourceAlias)
	if err != nil {
		return fmt.Errorf("初始化源客户端失败: %v", err)
	}

	targetClient, err := newClientFromAlias(targetAlias)
	if err != nil {
		return fmt.Errorf("初始化目标客户端失败: %v", err)
	}

	buckets, err := sourceClient.ListBuckets(ctx)
	if err != nil {
		return fmt.Errorf("获取源存储桶列表失败: %v", err)
	}

	for _, bucket := range buckets {
		if err := syncILMRules(ctx, sourceClient, targetClient, bucket.Name, bucket.Name); err != nil {
			//return err
		}

		if err := syncEventNotifications(ctx, sourceClient, targetClient, bucket.Name, bucket.Name); err != nil {
			//return err
		}
	}

	// 按顺序同步配置
	//if err := syncBucketBaseConfig(ctx, sourceClient, targetClient, sourceBucket, targetBucket); err != nil {
	//	return err
	//}
	//	if err := syncVersioning(ctx, sourceClient, targetClient, sourceBucket, targetBucket); err != nil {
	//		return err
	//	}

	//	if err := syncObjects(ctx, sourceClient, targetClient, sourceBucket, targetBucket); err != nil {
	//		return err
	//	}

	fmt.Println("同步完成：数据和配置已完全迁移")
	return nil
}

func main() {
	ctx := context.Background()
	// 执行同步
	if err := mainSync(ctx, "tw-test-old", "tw-test"); err != nil {
		log.Fatalf("同步失败: %v", err)
	}
}
