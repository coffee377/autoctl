package token

import (
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	redisOptions := redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "redis!@@&",      // Redis服务器密码
		DB:       0,                // Redis数据库索引
	}
	jwtGenerator := NewJWTGenerator(WithJinQi(), WithRedis(redisOptions))
	session := NewJinQiSession(WithAccount("coffee377"), WithFixSession(), WithExpire(time.Hour*24*7))
	token, err := jwtGenerator.Generate(session)
	t.Log(token)
	assert.NotNil(t, token)
	assert.Nil(t, err)
}
