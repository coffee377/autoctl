package process

import (
	"context"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
)

type InstanceMessageHandler func(ctx context.Context, header event.EventHeader, message *InstanceMessage) error
type TaskMessageHandler func(ctx context.Context, header event.EventHeader, message *TaskMessage) error

type EventFrameHandler struct {
	*event.DefaultEventFrameHandler

	instanceMessageHandler InstanceMessageHandler
	taskMessageHandler     TaskMessageHandler
}

func NewProcessEventFrameHandler() *EventFrameHandler {
	handler := &EventFrameHandler{}
	defaultHandler := handler.defaultHandler()
	handler.DefaultEventFrameHandler = event.NewDefaultEventFrameHandler(defaultHandler)
	return handler
}

func (efh *EventFrameHandler) SetInstanceMessageHandler(handler InstanceMessageHandler) {
	if handler != nil {
		efh.instanceMessageHandler = handler
	}
}

func (efh *EventFrameHandler) SetTaskMessageHandler(handler TaskMessageHandler) {
	if handler != nil {
		efh.taskMessageHandler = handler
	}
}

func (efh *EventFrameHandler) defaultHandler() event.IEventHandler {
	return func(c context.Context, header *event.EventHeader, rawData []byte) (event.EventProcessStatusType, error) {
		switch header.EventType {
		case "bpms_instance_change":
			instanceMessage := InstanceMessage{}
			if err := instanceMessage.UnmarshalJSON(rawData); err != nil {
				return event.EventProcessStatusKLater, err
			}
			if efh.instanceMessageHandler != nil {
				if err := efh.instanceMessageHandler(c, *header, &instanceMessage); err != nil {
					return event.EventProcessStatusKLater, err
				}
			}
		case "bpms_task_change":
			taskMessage := TaskMessage{}
			if err := taskMessage.UnmarshalJSON(rawData); err != nil {
				return event.EventProcessStatusKLater, err
			}
			if efh.taskMessageHandler != nil {
				if err := efh.taskMessageHandler(c, *header, &taskMessage); err != nil {
					return event.EventProcessStatusKLater, err
				}
			}
		}
		return event.EventProcessStatusKSuccess, nil

	}
}
