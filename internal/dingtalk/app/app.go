package app

import (
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	oauth21 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/coffee377/autoctl/internal/dingtalk"
)

type App interface {
	dingtalk.Namespace
	dingtalk.AccessToken

	GetID() string
	GetName() string

	GetClientID() string
	GetClientSecret() string

	GetAgentId() *string
	GetRobotCode() *string
}

func New(namespace, id string, options ...Option) App {
	application := &app{
		namespace:   namespace,
		id:          id,
		cachePrefix: "dingtalk",
	}
	for _, option := range options {
		option(application)
	}
	return application
}

type app struct {
	namespace string
	id        string // 应用ID

	clientId     string // 客户端 ID
	clientSecret string // 客户端密钥

	name      *string // 应用名称
	agentId   *string // 应用AgentId
	robotCode *string // 应用机器人编码

	cachePrefix          string
	cacheBeforeTokenHook func(context.Context) (string, bool)
	cacheAfterTokenHook  func(context.Context, string)
}

func (a *app) GetNamespaceName() string {
	return a.namespace
}

func (a *app) GetAccessToken(ctx context.Context) string {
	if a.cacheBeforeTokenHook != nil {
		if val, ok := a.cacheBeforeTokenHook(ctx); ok {
			return val
		}
	}

	config := new(openapi.Config)
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, _ := oauth21.NewClient(config)

	request := new(oauth21.GetAccessTokenRequest)
	request.SetAppKey(a.clientId)
	request.SetAppSecret(a.clientSecret)
	response, _ := client.GetAccessToken(request)
	accessToken := response.Body.AccessToken

	if a.cacheAfterTokenHook != nil {
		a.cacheAfterTokenHook(ctx, *accessToken)
	}

	return *accessToken
}

func (a *app) GetID() string {
	return a.id
}

func (a *app) GetName() string {
	return *a.name
}

func (a *app) GetClientID() string {
	return a.clientId
}

func (a *app) GetClientSecret() string {
	return a.clientSecret
}

func (a *app) GetAgentId() *string {
	return a.agentId
}

func (a *app) GetRobotCode() *string {
	return a.robotCode
}
