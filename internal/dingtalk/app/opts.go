package app

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Option func(*app)

func WithName(name string) Option {
	return func(o *app) {
		o.name = &name
	}
}

func WithClient(id, secret string) Option {
	return func(o *app) {
		o.clientId = id
		o.clientSecret = secret
	}
}

func WithAgent(id string) Option {
	return func(o *app) {
		o.agentId = &id
	}
}

func WithRobot(code string) Option {
	return func(o *app) {
		o.robotCode = &code
	}
}

func WithCachePrefix(prefix string) Option {
	return func(a *app) {
		a.cachePrefix = prefix
	}
}

func WithRedis(redisOpts redis.Options) Option {
	redisCli := redis.NewClient(&redisOpts)
	return func(o *app) {
		key := fmt.Sprintf("%s:%s:%s", o.cachePrefix, o.namespace, o.id)
		o.cacheBeforeTokenHook = func(ctx context.Context) (string, bool) {
			cmd := redisCli.Get(ctx, key)
			if cmd.Err() == nil {
				return cmd.Val(), true
			}
			return "", false
		}
		o.cacheAfterTokenHook = func(ctx context.Context, token string) {
			if token != "" {
				redisCli.Set(ctx, key, token, 2*time.Hour-5*time.Second)
			}
		}
	}
}
