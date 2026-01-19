package es

import (
	"cds/bid/ds"
	"cds/dingtalk/app"
	"cds/dingtalk/oa"

	"github.com/redis/go-redis/v9"
)

type SubscriptionOption func(sub *subscription)

func WithClient(id, secret string) SubscriptionOption {
	return func(sub *subscription) {
		sub.clientId = id
		sub.clientSecret = secret
	}
}

func WithRedis(options redis.Options) SubscriptionOption {
	return func(sub *subscription) {
		sub.rdb = redis.NewClient(&options)
	}
}

func WithApproval() SubscriptionOption {
	return func(sub *subscription) {
		approval, _ := oa.New(app.New("a57e9681-79cb-4242-96df-952be2dc3af7", app.WithRedis()))
		sub.approval = approval
	}
}

func WithEnt() SubscriptionOption {
	return func(sub *subscription) {
		mysqlClient, _ := ds.Mysql()
		sub.entClient = mysqlClient
	}
}

func WithProcessInstanceEvent(fn InstanceMessageHandler) SubscriptionOption {
	return func(sub *subscription) {
		if fn != nil {
			sub.instanceMessageHandler = fn
		}
	}
}

func WithProcessTaskEvent(fn TaskMessageHandler) SubscriptionOption {
	return func(sub *subscription) {
		if fn != nil {
			sub.taskMessageHandler = fn
		}
	}
}
