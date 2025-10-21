package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
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

type (
	User struct {
		UserId    string `json:"userid"`
		UnionId   string `json:"unionid"`
		RealName  string `json:"name"`
		JobNumber string `json:"job_number"`
		Email     string `json:"email"`
		Mobile    string `json:"mobile"`
		Avatar    string `json:"avatar"`
		Active    bool   `json:"active"`
	}

	Result[T any] struct {
		Code    int    `json:"errcode"`
		Message string `json:"errmsg"`
		Data    T      `json:"result"`
	}

	UserHook func(userId string) (*User, bool)
)

type App interface {
	Namespace
	AccessToken

	GetID() string
	GetName() string

	GetClientID() string
	GetClientSecret() string

	GetAgentId() string
	GetRobotCode() string

	GetUser(userId string) (*User, error)
	GetUserHook() UserHook
}

func (a *application) GetUserHook() UserHook {
	return func(userId string) (*User, bool) {
		if user, err := a.GetUser(userId); err == nil {
			return user, true
		}
		return nil, false
	}
}

func (a *application) GetUser(userId string) (*User, error) {
	params := url.Values{}
	params.Add("access_token", a.GetAccessToken())
	params.Add("userid", userId)
	resp, err := http.Get(fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/user/get?%s", params.Encode()))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body) // 确保关闭响应体，避免资源泄漏

	// 2. 读取响应体为字节流
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res Result[User]
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("%s", res.Message)
	}
	return &res.Data, nil
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
	confDir   string // 配置文件目录
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
