package redis

import "github.com/redis/go-redis/v9"

type Cache interface {
	Ping() bool
}

type cacheImpl struct {
	Cache
	rdb *redis.Client
}

func NewRedisCache() Cache {
	options := &redis.Options{}
	return &cacheImpl{
		rdb: redis.NewClient(options),
	}
}
