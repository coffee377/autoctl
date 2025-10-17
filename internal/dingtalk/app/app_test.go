package app

import (
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestNormalApp(t *testing.T) {
	a := New("118447d2-1c73-486f-8058-7daa046c9577")
	assert.NotNil(t, a)
	assert.Equal(t, "118447d2-1c73-486f-8058-7daa046c9577", a.GetID())
}

func TestAppWithName(t *testing.T) {
	a := New("118447d2-1c73-486f-8058-7daa046c9577", WithName("代码工匠实验室"))
	assert.Equal(t, "118447d2-1c73-486f-8058-7daa046c9577", a.GetID())
	assert.Equal(t, "代码工匠实验室", a.GetName())
}

func TestAppWithClient(t *testing.T) {
	a := New("118447d2-1c73-486f-8058-7daa046c9577", WithName("代码工匠实验室"),
		WithClient("1", "2"))
	assert.Equal(t, "118447d2-1c73-486f-8058-7daa046c9577", a.GetID())
	assert.Equal(t, "代码工匠实验室", a.GetName())
	assert.Equal(t, "1", a.GetClientID())
	assert.Equal(t, "2", a.GetClientSecret())
}

func TestAppOther(t *testing.T) {
	a := New("118447d2-1c73-486f-8058-7daa046c9577", WithName("代码工匠实验室"),
		WithNamespace("ccl"), WithClient("1", "2"),
		WithAgent("000"), WithRobot("robot"))
	assert.Equal(t, "118447d2-1c73-486f-8058-7daa046c9577", a.GetID())
	assert.Equal(t, "代码工匠实验室", a.GetName())
	assert.Equal(t, "ccl", a.GetNamespaceName())
	assert.Equal(t, "1", a.GetClientID())
	assert.Equal(t, "2", a.GetClientSecret())
	assert.Equal(t, "000", a.GetAgentId())
	assert.Equal(t, "robot", a.GetRobotCode())
}

func TestGetAccessToken(t *testing.T) {
	ap := New("118447d2-1c73-486f-8058-7daa046c9577",
		WithName("代码工匠实验室-监控平台-无缓存测试"),
	)
	token := ap.GetAccessToken()
	assert.NotNil(t, token)
}

func TestGetAccessTokenWithRedis(t *testing.T) {
	ap := New("118447d2-1c73-486f-8058-7daa046c9577",
		WithName("代码工匠实验室-监控平台-Redis缓存测试"),
		WithRedis(redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "redis!@@&",
			DB:       0,
		}),
		WithCachePrefix("dingtalk"),
	)
	token := ap.GetAccessToken()
	assert.NotNil(t, token)
}
