package dingtalk

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	oauth21 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/coffee377/autoctl/log"
)

// GetAccessToken 获取企业内部应用的accessToken
// https://open.dingtalk.com/document/orgapp-server/obtain-the-access_token-of-an-internal-app
func GetAccessToken() {
	c := new(openapi.Config)
	c.SetProtocol("https")
	c.SetRegionId("central")
	//c.SetEndpoint("https://oapi.dingtalk.com/gettoken")
	client, _ := oauth21.NewClient(c)

	//config := &openapi.Config{}
	//config.Protocol = tea.String("https")
	//config.RegionId = tea.String("central")
	//_result = &dingtalkoauth2_1_0.Client{}
	//_result, _err = dingtalkoauth2_1_0.NewClient(config)

	request := new(oauth21.GetAccessTokenRequest)
	request.SetAppKey("dingopfniakkw72klkjv")
	request.SetAppSecret("6Il0DuPZPPIr-OG03uMrnqDNu_o03tpIkK03ScpuEPP6NAw7J52D0LWPvTjRf4BR")
	token, err := client.GetAccessToken(request)
	var result *string
	result = token.Body.AccessToken
	log.Debug("AccessToken: %s", *result)
	if err != nil {
		return
	}

}
