package es

import (
	"cds/dingtalk/es/process"
	"context"
	"fmt"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/card"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/plugin"
)

type Subscription interface {
	Run(ctx context.Context) error
}

func DingTalkEventSubscription(options ...SubscriptionOption) Subscription {
	s := &subscription{
		processEventFrameHandler: process.NewProcessEventFrameHandler(),
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

type subscription struct {
	clientId     string
	clientSecret string

	processEventFrameHandler *process.EventFrameHandler
}

func (s *subscription) Run(ctx context.Context) error {

	logger.SetLogger(logger.NewStdTestLogger())

	cli := client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(s.clientId, s.clientSecret)))

	// 注册事件类型的处理函数
	if s.processEventFrameHandler != nil {
		cli.RegisterAllEventRouter(s.processEventFrameHandler.OnEventReceived)
	}
	// 注册机器人callback类型的处理函数
	cli.RegisterChatBotCallbackRouter(OnChatBotMessageReceived)
	//注册AI插件的处理函数
	cli.RegisterPluginCallbackRouter(OnPluginMessageReceived)
	// 注册互动卡片类型的处理函数
	cli.RegisterCardCallbackRouter(OnCardCallbackReceived)

	err := cli.Start(ctx)
	if err != nil {
		return err
	}
	defer cli.Close()

	select {}
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
