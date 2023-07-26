package dingtalk

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkim10 "github.com/alibabacloud-go/dingtalk/im_1_0"
	oauth21 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/coffee377/autoctl/lib/log"
)

// GetAccessToken 获取企业内部应用的accessToken
// https://open.dingtalk.com/document/orgapp-server/obtain-the-access_token-of-an-internal-app
func GetAccessToken() {
	config := new(openapi.Config)
	//config := &openapi.Config{}
	//config.SetProtocol()
	config.SetProtocol("https")
	config.SetRegionId("central")
	//c.SetEndpoint("https://oapi.dingtalk.com/gettoken")
	client, _ := oauth21.NewClient(config)

	//config := &openapi.Config{}
	//config.Protocol = tea.String("https")
	//config.RegionId = tea.String("central")
	//_result = &dingtalkoauth2_1_0.Client{}
	//_result, _err = dingtalkoauth2_1_0.NewClient(config)

	request := new(oauth21.GetAccessTokenRequest)
	request.SetAppKey("dingopfniakkw72klkjv")
	request.SetAppSecret("6Il0DuPZPPIr-OG03uMrnqDNu_o03tpIkK03ScpuEPP6NAw7J52D0LWPvTjRf4BR")
	response, err := client.GetAccessToken(request)
	var accessToken *string
	accessToken = response.Body.AccessToken
	log.Debug("AccessToken: %s", *accessToken)
	if err != nil {
		return
	}

	imClient, _ := dingtalkim10.NewClient(config)

	interactiveCardCreateInstanceHeaders := &dingtalkim10.InteractiveCardCreateInstanceHeaders{}
	interactiveCardCreateInstanceHeaders.SetXAcsDingtalkAccessToken(*accessToken)

	interactiveCardCreateInstanceRequest := &dingtalkim10.InteractiveCardCreateInstanceRequest{}
	interactiveCardCreateInstanceRequest.SetCardTemplateId("") // 互动卡片模板 id

	interactiveCardCreateInstanceRequest.SetConversationType(1)    // 会话类型 0：单聊 1：群聊
	interactiveCardCreateInstanceRequest.SetOpenConversationId("") // 群会话 id，群聊时设置
	interactiveCardCreateInstanceRequest.SetOutTrackId("")         // 唯一标识一张卡片的外部ID,可用于更新或重复发送同一卡片
	//interactiveCardCreateInstanceRequest.SetReceiverUserIdList(tea.StringSlice([]string{"", ""}))

	// 卡片模板数据
	cardDataCardMediaIdParamMap := map[string]*string{
		"key": tea.String("sfrtxxxx"),
	}
	cardDataCardParamMap := map[string]*string{
		"key": tea.String("afxxxx"),
	}
	cardData := &dingtalkim10.InteractiveCardCreateInstanceRequestCardData{
		CardParamMap:        cardDataCardParamMap,
		CardMediaIdParamMap: cardDataCardMediaIdParamMap,
	}
	interactiveCardCreateInstanceRequest.SetCardData(cardData)

	// 指定用户可见按钮列表
	privateDataValueKeyCardMediaIdParamMap := map[string]*string{
		"key": tea.String("xxxx"),
	}
	privateDataValueKeyCardParamMap := map[string]*string{
		"key": tea.String("wwhtxxxx"),
	}
	privateDataValueKey := &dingtalkim10.PrivateDataValue{
		CardParamMap:        privateDataValueKeyCardParamMap,
		CardMediaIdParamMap: privateDataValueKeyCardMediaIdParamMap,
	}
	privateData := map[string]*dingtalkim10.PrivateDataValue{
		"privateDataValueKey": privateDataValueKey,
	}
	interactiveCardCreateInstanceRequest.SetPrivateData(privateData)

	//interactiveCardCreateInstanceHeaders.XAcsDingtalkAccessToken = result
	imClient.InteractiveCardCreateInstance(interactiveCardCreateInstanceRequest)
}
