package dingtalk

import (
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	card10 "github.com/alibabacloud-go/dingtalk/card_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/coffee377/autoctl/pkg/log"
)

func CreateClient() (_result *card10.Client, _err error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &card10.Client{}
	_result, _err = card10.NewClient(config)
	return _result, _err
}

type Card struct {
	app    *App
	client *card10.Client
}

func NewCard(app *App) (*Card, error) {
	config := new(openapi.Config)
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, err := card10.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Card{app: app, client: client}, nil
}

//type ICard interface {
//	CreateCard(card *card10.CreateCardRequest) (card10.CreateCardResponse, error)
//}

func (c Card) Create(templateId string) (string, error) {
	accessToken := c.app.GetAccessToken()
	headers := &card10.CreateCardHeaders{
		XAcsDingtalkAccessToken: tea.String(accessToken),
	}
	request := &card10.CreateCardRequest{}

	createdAt := time.Now()
	outTrackId := createdAt.Format("20060102150405")
	log.Info("outTrackId: %s", outTrackId)

	request.SetCardTemplateId(templateId)
	request.SetOutTrackId(outTrackId)
	request.SetCallbackType("STREAM")
	request.SetCardData(&card10.CreateCardRequestCardData{
		CardParamMap: map[string]*string{
			"title":           tea.String("CI/CD"),
			"abstract":        tea.String("CI/CD投出提示"),
			"env_name":        tea.String("测试环境"),
			"env_color":       tea.String("orange"),
			"stage":           tea.String("Maven 构建"),
			"consumeTime":     tea.String("8h8min"),
			"status":          tea.String("出差费用"),
			"createdAt":       tea.String(createdAt.Format(time.DateTime)),
			"sys_lastMessage": tea.String("审批"),
		},
	})

	response, err := c.client.CreateCardWithOptions(request, headers, &util.RuntimeOptions{})
	if err != nil {
		return "", err
	}
	return *response.Body.Result, nil
}

func (c Card) CreateAndDeliver(templateId string) (*card10.CreateAndDeliverResponseBodyResult, error) {
	accessToken := c.app.GetAccessToken()
	headers := &card10.CreateAndDeliverHeaders{
		XAcsDingtalkAccessToken: tea.String(accessToken),
	}
	request := &card10.CreateAndDeliverRequest{}

	createdAt := time.Now()
	outTrackId := createdAt.Format("20060102150405")
	log.Info("outTrackId: %s", outTrackId)

	request.SetCardTemplateId(templateId)
	request.SetOutTrackId(outTrackId)
	request.SetCallbackType("STREAM")
	request.SetCardData(&card10.CreateAndDeliverRequestCardData{
		CardParamMap: map[string]*string{
			"title":               tea.String("CI/CD"),
			"abstract":            tea.String("CI/CD透出提示"),
			"sys_lastMessageI18n": tea.String("{\"zh_CN\":\"蚂蚁分工\",\"zh_TW\":\"螞蟻分工\",\"zh_HK\":\"螞蟻分工\",\"ja_JP\":\"アリの分業\",\"en_US\":\"Ant division of labor\"}"),
			//"env":                 tea.String("{\"name\": \"测试环境\",\"color\": \"orange\"}"),
			"env":         tea.String("{\"name\": \"预生产环境\",\"color\": \"green\"}"),
			"stage":       tea.String("Maven 构建"),
			"consumeTime": tea.String("8h8min"),
			"status":      tea.String("2"),
			"createdAt":   tea.String(createdAt.Format(time.DateTime)),
			"buildable":   tea.String("true"),
			"deployable":  tea.String("false"),
			"reversible":  tea.String("false"),
		},
	})
	request.SetPrivateData(map[string]*card10.PrivateDataValue{
		//"02140408367343": {
		//	CardParamMap: map[string]*string{
		//		//"outTrackId": tea.String(outTrackId),
		//	},
		//},
	})

	//  1. IM群聊
	//request.SetOpenSpaceId("dtv1.card//im_group.cidXdtJXrL/VA2X4/C/MQA/6g==")
	//request.SetImGroupOpenSpaceModel(&card10.CreateAndDeliverRequestImGroupOpenSpaceModel{
	//	SupportForward: tea.Bool(false),
	//})
	//request.SetImGroupOpenDeliverModel(&card10.CreateAndDeliverRequestImGroupOpenDeliverModel{
	//	RobotCode:  tea.String(c.app.RobotCode),
	//	Recipients: tea.StringSlice([]string{"02140408367343", "016729230300631918553"}),
	//	AtUserIds: map[string]*string{
	//		"02140408367343": tea.String("test"),
	//	},
	//})

	// 2. 群置顶卡片（吊顶）
	//request.SetOpenSpaceId("dtv1.card//ONE_BOX.cidXdtJXrL/VA2X4/C/MQA/6g==")
	//request.SetTopOpenSpaceModel(&card10.CreateAndDeliverRequestTopOpenSpaceModel{
	//	SpaceType: tea.String("ONE_BOX"),
	//})
	//request.SetTopOpenDeliverModel(&card10.CreateAndDeliverRequestTopOpenDeliverModel{
	//	ExpiredTimeMillis: tea.Int64(time.Now().Add(time.Second * 10).UnixMilli()),
	//	UserIds:           tea.StringSlice([]string{"02140408367343"}),
	//})

	// 3. IM机器人单聊
	request.SetOpenSpaceId("dtv1.card//im_robot.02140408367343")
	request.SetImRobotOpenSpaceModel(&card10.CreateAndDeliverRequestImRobotOpenSpaceModel{
		SupportForward: tea.Bool(false),
	})
	request.SetImRobotOpenDeliverModel(&card10.CreateAndDeliverRequestImRobotOpenDeliverModel{
		RobotCode: tea.String(c.app.RobotCode),
		SpaceType: tea.String("IM_ROBOT"),
	})

	// IM_SINGLE
	//request.SetOpenSpaceId("dtv1.card//im_single.cidXdtJXrL/VA2X4/C/MQA/6g==")
	//request.SetImSingleOpenSpaceModel(&card10.CreateAndDeliverRequestImSingleOpenSpaceModel{
	//	SupportForward: tea.Bool(false),
	//})
	//request.SetImSingleOpenDeliverModel(&card10.CreateAndDeliverRequestImSingleOpenDeliverModel{
	//	AtUserIds: map[string]*string{
	//		"02140408367343": tea.String("wyj"),
	//	},
	//})

	// 5. 协作 ???
	//request.SetOpenSpaceId("dtv1.card//COOPERATON_FEED.02140408367343")
	//request.SetCoFeedOpenSpaceModel(&card10.CreateAndDeliverRequestCoFeedOpenSpaceModel{
	//	CoolAppCode: tea.String("cool"),
	//	Title:       tea.String("协作卡片测试"),
	//})
	//request.SetCoFeedOpenDeliverModel(&card10.CreateAndDeliverRequestCoFeedOpenDeliverModel{
	//	BizTag:      tea.String("3333333333"),
	//	GmtTimeLine: tea.Int64(0),
	//})

	response, err := c.client.CreateAndDeliverWithOptions(request, headers, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	return response.Body.Result, nil
}

func (c Card) Deliver(outTrackId string) ([]*card10.DeliverCardResponseBodyResult, error) {
	accessToken := c.app.GetAccessToken()
	headers := &card10.DeliverCardHeaders{
		XAcsDingtalkAccessToken: tea.String(accessToken),
	}
	request := &card10.DeliverCardRequest{}
	request.SetOutTrackId(outTrackId)
	request.SetOpenSpaceId("dtv1.card//IM_GROUP.cidXdtJXrL/VA2X4/C/MQA/6g==")
	request.SetImGroupOpenDeliverModel(&card10.DeliverCardRequestImGroupOpenDeliverModel{
		RobotCode:  tea.String(c.app.RobotCode),
		Recipients: tea.StringSlice([]string{"02140408367343"}),
	})
	result, err := c.client.DeliverCardWithOptions(request, headers, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	return result.Body.Result, nil
}

func (c Card) Update() error {
	return nil
}

func (c Card) AddOrUpdateSpaces() error {
	return nil
}
