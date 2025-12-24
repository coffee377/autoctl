package es

import (
	"context"
	"encoding/json"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
	"github.com/redis/go-redis/v9"
)

type RedisEventFrameHandler struct {
	defaultHandler EventHandler
}

func NewRedisEventFrameHandler(rdb *redis.Client) *RedisEventFrameHandler {
	return &RedisEventFrameHandler{
		defaultHandler: RedisEventHandler(rdb),
	}
}

func (h *RedisEventFrameHandler) OnEventReceived(ctx context.Context, df *payload.DataFrame) (*payload.DataFrameResponse, error) {
	eventHeader := event.NewEventHeaderFromDataFrame(df)
	if df == nil || h.defaultHandler == nil {
		logger.GetLogger().Warningf("No event handler found, drop this event. eventType=[%s], eventId=[%s], eventCorpId=[%s]",
			eventHeader.EventType, eventHeader.EventId, eventHeader.EventCorpId)
		return nil, nil
	}

	ret, err := h.defaultHandler(ctx, *eventHeader, []byte(df.Data), *df)
	if err != nil {
		logger.GetLogger().Errorf("Event handler process error. eventType=[%s], eventId=[%s], eventCorpId=[%s] err=[%s]",
			eventHeader.EventType, eventHeader.EventId, eventHeader.EventCorpId, err)
		ret = event.EventProcessStatusKLater
	}

	result := event.NewEventProcessResultSuccess()
	code := payload.DataFrameResponseStatusCodeKOK
	if ret != event.EventProcessStatusKSuccess {
		code = payload.DataFrameResponseStatusCodeKInternalError
		result = event.NewEventProcessResultLater()
	}

	resultStr, _ := json.Marshal(result)
	frameResp := &payload.DataFrameResponse{
		Code: code,
		Headers: payload.DataFrameHeader{
			payload.DataFrameHeaderKContentType: payload.DataFrameContentTypeKJson,
			payload.DataFrameHeaderKMessageId:   df.GetMessageId(),
		},
		Message: "ok",
		Data:    string(resultStr),
	}

	return frameResp, nil
}
