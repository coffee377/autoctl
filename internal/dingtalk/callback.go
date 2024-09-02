package dingtalk

import (
	"context"
	"fmt"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/card"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/plugin"
)

const (
	clientId     = "dingybihm3fg4sjh3dtx"
	clientSecret = "smpvcY639CMUdAfmWOoyIImFCdD0woA09cMp7S5AsAQZGki6XFUUVrp0XCUCE-N2"
)

func SimpleReplyCard(ctx context.Context, sessionWebhook string, r *chatbot.ChatbotReplier, title, content []byte) error {
	requestBody := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": string(title),
			"text":  string(content),
		},
	}
	return r.ReplyMessage(ctx, sessionWebhook, requestBody)
}

// OnChatBotMessageReceived 简单的应答机器人实现
func OnChatBotMessageReceived(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	replyMsg := []byte(fmt.Sprintf("msg received: [%s]", data.Text.Content))

	chatbotReplier := chatbot.NewChatbotReplier()
	if err := chatbotReplier.SimpleReplyText(ctx, data.SessionWebhook, replyMsg); err != nil {
		return nil, err
	}
	//if err := chatbotReplier.SimpleReplyMarkdown(ctx, data.SessionWebhook, []byte("Markdown消息"), replyMsg); err != nil {
	//	return nil, err
	//}

	//SimpleReplyCard(ctx, data.SessionWebhook, chatbotReplier, []byte("Markdown消息"), replyMsg)

	return replyMsg, nil
}

// OnPluginMessageReceived 简单的插件处理实现
func OnPluginMessageReceived(ctx context.Context, request *plugin.GraphRequest) (*plugin.GraphResponse, error) {
	response := &plugin.GraphResponse{
		Body: `{"text": "hello world", "content": [{"title": "1", "description": "2", "url":"https://www.zhihu.com/question/626551401"},{"title": "2", "description": "2", "url":"https://www.zhihu.com/question/626551401"}]}`,
	}
	return response, nil
}

// OnEventReceived 事件处理
// https://open.dingtalk.com/document/orgapp/event-subscription-overview#8dcdbb72adhxy
func OnEventReceived(ctx context.Context, df *payload.DataFrame) (frameResp *payload.DataFrameResponse, err error) {
	eventHeader := event.NewEventHeaderFromDataFrame(df)

	logger.GetLogger().Infof("received event, eventId=[%s] eventBornTime=[%d] eventCorpId=[%s] eventType=[%s] eventUnifiedAppId=[%s] data=[%s]",
		eventHeader.EventId,
		eventHeader.EventBornTime,
		eventHeader.EventCorpId,
		eventHeader.EventType,
		eventHeader.EventUnifiedAppId,
		df.Data)

	frameResp = payload.NewSuccessDataFrameResponse()
	if err := frameResp.SetJson(event.NewEventProcessResultSuccess()); err != nil {
		return nil, err
	}

	return
}

// OnCardCallbackReceived 卡片回调处理
func OnCardCallbackReceived(ctx context.Context, request *card.CardRequest) (*card.CardResponse, error) {
	logger.GetLogger().Infof("receive card data: %v", request)
	action := request.CardActionData.CardPrivateData.Params["action"]
	logger.GetLogger().Infof("action: %s", action)

	cardParamMap := map[string]string{}
	switch action {
	case "build":
		cardParamMap["buildable"] = "false"
		cardParamMap["deployable"] = "true"
		cardParamMap["reversible"] = "false"
	case "deploy":
		cardParamMap["buildable"] = "false"
		cardParamMap["deployable"] = "false"
		cardParamMap["reversible"] = "true"
	case "rollback":
		cardParamMap["buildable"] = "false"
		cardParamMap["deployable"] = "false"
		cardParamMap["reversible"] = "false"
	}

	response := &card.CardResponse{
		CardUpdateOptions: &card.CardUpdateOptions{
			UpdateCardDataByKey:    false,
			UpdatePrivateDataByKey: true,
		},
		CardData: &card.CardDataDto{
			CardParamMap: map[string]string{},
		},
		UserPrivateData: &card.CardDataDto{
			CardParamMap: cardParamMap,
		},
	}
	return response, nil
}

func Start() {
	logger.SetLogger(logger.NewStdTestLogger())
	cli := client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(clientId, clientSecret)))

	// 注册事件类型的处理函数
	cli.RegisterAllEventRouter(OnEventReceived)
	// 注册callback类型的处理函数
	cli.RegisterChatBotCallbackRouter(OnChatBotMessageReceived)
	// 注册AI插件的处理函数
	cli.RegisterPluginCallbackRouter(OnPluginMessageReceived)
	// 注册互动卡片类型的处理函数
	cli.RegisterCardCallbackRouter(OnCardCallbackReceived)

	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	select {}
}
