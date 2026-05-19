package cmd

import (
	"fmt"

	"github.com/coffee377/autoctl/configs"
	"github.com/coffee377/autoctl/internal/token"
	"github.com/redis/go-redis/v9"
)

type TokenExecutor struct {
	config *configs.Config
}

func NewTokenExecutor(cfg *configs.Config) *TokenExecutor {
	return &TokenExecutor{config: cfg}
}

func (e *TokenExecutor) Execute() (string, error) {
	redisOptions := e.buildRedisOptions()

	jwtOpts := []token.JWTGeneratorOptions{
		token.WithRedis(redisOptions),
		token.WithJinQi(),
	}

	jwtGenerator := token.NewJWTGenerator(jwtOpts...)

	sessionOpts := e.buildSessionOptions()

	session := token.NewJinQiSession(sessionOpts...)
	return jwtGenerator.Generate(session)
}

func (e *TokenExecutor) buildRedisOptions() redis.Options {
	opts := redis.Options{}

	if e.config.Redis.Address != "" {
		opts.Addr = e.config.Redis.Address
	} else if e.config.Redis.Host != "" {
		port := 6379
		if e.config.Redis.Port > 0 {
			port = e.config.Redis.Port
		}
		opts.Addr = fmt.Sprintf("%s:%d", e.config.Redis.Host, port)
	}

	if e.config.Redis.Password != "" {
		opts.Password = e.config.Redis.Password
	}

	if e.config.Redis.DB > 0 {
		opts.DB = e.config.Redis.DB
	}

	return opts
}

func (e *TokenExecutor) buildSessionOptions() []token.SessionOptions {
	opts := []token.SessionOptions{
		token.WithAccount(e.config.Token.Account),
		token.WithExpire(e.config.Token.Expire),
	}

	if e.config.Token.SessionFixed {
		opts = append(opts, token.WithFixSession())
	}

	return opts
}