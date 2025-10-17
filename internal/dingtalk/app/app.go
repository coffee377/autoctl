package app

import (
	"context"
	"errors"
	"log/slog"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	oauth21 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
)

type Namespace interface {
	GetNamespaceName() string // 获取命名空间名称
}

type AccessToken interface {
	GetAccessToken() string
}

type App interface {
	Namespace
	AccessToken

	GetID() string
	GetName() string

	GetClientID() string
	GetClientSecret() string

	GetAgentId() string
	GetRobotCode() string
}

func New(appId string, options ...Option) App {
	configurations := ReadConfiguration()
	var config *Configuration
	for _, app := range configurations.App {
		if app.Id == appId {
			config = &app
			break
		}
	}

	app := &application{
		cachePrefix: "dingtalk",
	}

	if config != nil {
		app.namespace = config.Namespace
		app.id = config.Id
		app.name = config.Name
		app.clientId = config.ClientId
		app.clientSecret = config.ClientSecret
		app.agentId = config.AgentId
		app.robotCode = config.RobotCode
	}

	for _, option := range options {
		option(app)
	}
	err := validateApp(app)
	if err != nil {
		slog.Error("validate app failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	return app
}

func validateApp(application *application) error {
	if application.namespace == "" {
		return errors.New("app namespace is required")
	}
	if application.id == "" {
		return errors.New("app id is required")
	}
	if application.clientId == "" || application.clientSecret == "" {
		return errors.New("app clientId and clientSecret is required")
	}
	return nil
}

type application struct {
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

func (a *application) GetNamespaceName() string {
	return a.namespace
}

func (a *application) GetAccessToken() string {
	ctx := context.TODO()
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

func (a *application) GetID() string {
	return a.id
}

func (a *application) GetName() string {
	return *a.name
}

func (a *application) GetClientID() string {
	return a.clientId
}

func (a *application) GetClientSecret() string {
	return a.clientSecret
}

func (a *application) GetAgentId() string {
	return *a.agentId
}

func (a *application) GetRobotCode() string {
	return *a.robotCode
}
