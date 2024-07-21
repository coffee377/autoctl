package dingtalk

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	oauth21 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
)

//type EventSubscription struct {
//	aesKey string
//	token  string
//	url    string
//}

type Base struct {
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
	AppKey string
	/**
	 * 应用的密钥
	 */
	AppSecret string
}

type AccessToken interface {
	GetAccessToken() string
}

type App struct {
	Base
	//Primary           bool
	//eventSubscription EventSubscription
}

func (a *App) GetAccessToken() string {
	config := new(openapi.Config)
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, _ := oauth21.NewClient(config)

	request := new(oauth21.GetAccessTokenRequest)
	request.SetAppKey(a.AppKey)
	request.SetAppSecret(a.AppSecret)
	response, _ := client.GetAccessToken(request)
	accessToken := response.Body.AccessToken
	//log.Debug("AccessToken: %s", *accessToken)
	return *accessToken
}

//type Robot struct {
//	Base
//}
//
//type DingTalk struct {
//	CorpId string
//	Apps   []App
//}
//
//type IDingTalk interface {
//	GetApps() []App
//}
