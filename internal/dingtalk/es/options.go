package es

import (
	"cds/dingtalk/es/process"
)

type SubscriptionOption func(sub *subscription)

func WithClient(id, secret string) SubscriptionOption {
	return func(sub *subscription) {
		sub.clientId = id
		sub.clientSecret = secret
	}
}

func WithProcessInstanceEvent(handler process.InstanceMessageHandler) SubscriptionOption {
	// 商务部投标模块 项目付款 - 审批实例开始结束
	// /v1.0/event/bpms_instance_change/processCode/PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97/type/{start,finish,terminate,delete}
	//topic := fmt.Sprintf("/v1.0/event/bpms_instance_change/processCode/%s/type/*", processCode)
	return func(sub *subscription) {
		sub.processEventFrameHandler.SetInstanceMessageHandler(handler)
	}
}

func WithProcessTaskEvent(handler process.TaskMessageHandler) SubscriptionOption {
	// 商务部投标模块 项目付款 - 审批任务开始、结束、转交
	// /v1.0/event/bpms_task_change/processCode/PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97/type/{start,finish}
	//topic := fmt.Sprintf("/v1.0/event/bpms_task_change/processCode/%s/type/*", processCode)
	return func(sub *subscription) {
		//sub.events[topic] = handler
		sub.processEventFrameHandler.SetTaskMessageHandler(handler)
	}
}
