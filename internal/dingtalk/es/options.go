package es

import (
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
