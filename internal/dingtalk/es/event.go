package es

import (
	"context"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
)

type EventHandler func(c context.Context, header event.EventHeader, rawData []byte, df payload.DataFrame) (event.EventProcessStatusType, error)
