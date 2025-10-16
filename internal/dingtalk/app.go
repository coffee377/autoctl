package dingtalk

import (
	"context"
	"fmt"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	oauth21 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/coffee377/autoctl/pkg/log"
	"github.com/redis/go-redis/v9"
)

type App struct {
	redis *redis.Client

	Id           string  // 应用ID
	Name         *string // 应用名称
	AgentId      string  // 应用AgentId
	ClientID     string  // 客户端 ID
	ClientSecret string  // 客户端密钥

	RobotCode string // 应用机器人编码

	// 是否主应用
	Primary bool
}

// GetAccessToken 获取企业内部应用的accessToken
func (a *App) GetAccessToken() string {
	ctx := context.TODO()
	if val, ok := a.beforeAccessToken(ctx); ok {
		return val
	}

	config := new(openapi.Config)
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, _ := oauth21.NewClient(config)

	request := new(oauth21.GetAccessTokenRequest)
	request.SetAppKey(a.ClientID)
	request.SetAppSecret(a.ClientSecret)
	response, _ := client.GetAccessToken(request)
	accessToken := response.Body.AccessToken

	a.afterAccessToken(ctx, accessToken, time.Hour*2)

	return *accessToken
}

func (a *App) beforeAccessToken(ctx context.Context) (string, bool) {
	if a.redis == nil {
		a.redis = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "redis!@@&",
			DB:       0,
		})
	}

	cmd := a.redis.Get(ctx, a.key())
	if cmd.Err() == nil {
		return cmd.Val(), true
	}
	return "", false
}

func (a *App) afterAccessToken(ctx context.Context, token *string, duration time.Duration) {
	log.Info("AccessToken: %s", *token)
	a.redis.Set(ctx, a.key(), token, duration)
}

func (a *App) key() string {
	return fmt.Sprintf("dingtalk:%s:%s", "", a.Id)
}
