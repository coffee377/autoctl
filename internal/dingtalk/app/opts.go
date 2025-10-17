package app

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Option func(*application)

func WithNamespace(namespace string) Option {
	return func(o *application) {
		o.namespace = namespace
	}
}

func WithName(name string) Option {
	return func(o *application) {
		o.name = &name
	}
}

func WithClient(id, secret string) Option {
	return func(o *application) {
		o.clientId = id
		o.clientSecret = secret
	}
}

func WithAgent(id string) Option {
	return func(o *application) {
		o.agentId = &id
	}
}

func WithRobot(code string) Option {
	return func(o *application) {
		o.robotCode = &code
	}
}

func WithCachePrefix(prefix string) Option {
	return func(a *application) {
		a.cachePrefix = prefix
	}
}

func WithRedis(redisOpts ...redis.Options) Option {
	var redisCli *redis.Client
	if len(redisOpts) == 0 {
		redisCli = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "redis!@@&",
			DB:       0,
		})
	} else {
		redisCli = redis.NewClient(&redisOpts[0])
	}
	return func(o *application) {
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
