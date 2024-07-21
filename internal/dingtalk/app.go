package dingtalk

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	oauth21 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/coffee377/autoctl/pkg/log"
)

type AccessToken interface {
	GetAccessToken() string
}

type App struct {
	Id string
	/**
	 * 应用名称
	 */
	Name string
	/**
	 * 应用的唯一标识
	 */
	AgentId string
	/**
	 * 应用的 Key
	 */
	ClientKey string
	/**
	 * 应用的密钥
	 */
	ClientSecret string

	/**
	 * 应用机器人编码
	 */
	RobotCode string
	// 是否主应用
	Primary bool
}

// GetAccessToken 获取企业内部应用的accessToken
func (a *App) GetAccessToken() string {
	// todo 使用 redis 缓存 access token
	config := new(openapi.Config)
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, _ := oauth21.NewClient(config)

	request := new(oauth21.GetAccessTokenRequest)
	request.SetAppKey(a.ClientKey)
	request.SetAppSecret(a.ClientSecret)
	response, _ := client.GetAccessToken(request)
	accessToken := response.Body.AccessToken
	log.Info("AccessToken: %s", *accessToken)
	return *accessToken
}
