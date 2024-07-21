package dingtalk

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkim10 "github.com/alibabacloud-go/dingtalk/im_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/coffee377/autoctl/pkg/log"
	"time"
)

const (
	cardTemplateId = "49528a07-0818-4306-9066-4e9775744bf1.schema"
	robotCode      = "dingybihm3fg4sjh3dtx"
)

type Robot struct {
	app    *App
	client *dingtalkim10.Client
}

func NewRobot(app *App) (*Robot, error) {
	config := new(openapi.Config)
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, err := dingtalkim10.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Robot{app: app, client: client}, nil
}

type ChatType int

const (
	SingleChat ChatType = 0 // 单聊
	GroupChat  ChatType = 1 // 群聊
)

// SendCardMessage https://open.dingtalk.com/document/orgapp/send-interactive-dynamic-cards-1
func (r Robot) SendCardMessage(chatType ChatType) (string, error) {
	accessToken := r.app.GetAccessToken()
	headers := &dingtalkim10.SendInteractiveCardHeaders{}
	headers.SetXAcsDingtalkAccessToken(accessToken)

	request := &dingtalkim10.SendInteractiveCardRequest{}
	request.SetCardTemplateId(cardTemplateId) // 互动卡片模板 id
	createdAt := time.Now()
	outTrackId := createdAt.Format("20060102150405")
	log.Info("outTrackId: %s", outTrackId)

	request.SetConversationType(int32(chatType)) // 会话类型 0：单聊 1：群聊
	request.SetUserIdType(1)
	if chatType == GroupChat {
		// 研发中心 chatId => chatec160f5439b459a57c55391a21d7b27b , openConversationId => cidXdtJXrL/VA2X4/C/MQA/6g==
		request.SetOpenConversationId("cidXdtJXrL/VA2X4/C/MQA/6g==") // 群会话 id，群聊时设置
	}

	request.SetReceiverUserIdList(tea.StringSlice([]string{"02140408367343"})) // 接收者用户ID列表
	request.SetOutTrackId(outTrackId)                                          // 唯一标识一张卡片的外部ID,可用于更新或重复发送同一卡片
	if r.app.RobotCode != "" {
		request.SetRobotCode(r.app.RobotCode)
	}

	// 卡片模板数据
	cardData := &dingtalkim10.SendInteractiveCardRequestCardData{
		CardParamMap: map[string]*string{
			"title":       tea.String("朱小志提交的财务报销"),
			"type":        tea.String("差旅费"),
			"amount":      tea.String("1000元"),
			"reason":      tea.String("出差费用"),
			"lastMessage": tea.String("审批"),
			"created_at":  tea.String(createdAt.Format(time.DateTime)),
		},
		CardMediaIdParamMap: map[string]*string{},
	}
	request.SetCardData(cardData)

	// 指定用户私有数据列表
	privateData := map[string]*dingtalkim10.PrivateDataValue{
		"02140408367343": {
			CardParamMap: map[string]*string{
				"key": tea.String("withXXX"),
			},
			CardMediaIdParamMap: map[string]*string{
				"key": tea.String("xxx"),
			},
		},
	}
	request.SetPrivateData(privateData)

	instance, err := r.client.SendInteractiveCardWithOptions(request, headers, &util.RuntimeOptions{})
	if err != nil {
		return "", err
	}
	return *instance.Body.Result.ProcessQueryKey, nil
}

// UpdateCardMessage https://open.dingtalk.com/document/orgapp/update-dingtalk-interactive-cards
func (r Robot) UpdateCardMessage(outTrackId string) error {
	return nil
}
