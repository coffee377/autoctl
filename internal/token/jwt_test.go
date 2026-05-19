package token

import (
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	redisOptions := redis.Options{
		//Addr:     "localhost:6379", // Redis 服务器地址
		//Password: "redis!@@&",      // Redis 服务器密码
		//DB:       0,                // Redis 数据库索引

		Addr:     "192.168.44.82:6379",
		Password: "jqkj5350**)",
		DB:       5,
	}
	jwtGenerator := NewJWTGenerator(WithJinQi(), WithRedis(redisOptions))
	session := NewJinQiSession(WithAccount("coffee377"), WithFixSession(), WithExpire(time.Hour*24*7))
	token, err := jwtGenerator.Generate(session)
	t.Log(token)
	assert.NotNil(t, token)
	assert.Nil(t, err)
}
