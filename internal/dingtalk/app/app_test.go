package app

import (
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var (
	testClientId     = os.Getenv("APP_CLIENT_ID")
	testClientSecret = os.Getenv("APP_CLIENT_SECRET")
)

func TestNormalApp(t *testing.T) {
	a := New("ccl", "123456")
	assert.Equal(t, "ccl", a.GetNamespaceName())
	assert.Equal(t, "123456", a.GetID())
}

func TestAppWithName(t *testing.T) {
	a := New("ccl", "123456", WithName("代码工匠实验室"))
	assert.Equal(t, "ccl", a.GetNamespaceName())
	assert.Equal(t, "123456", a.GetID())
	assert.Equal(t, "代码工匠实验室", a.GetName())
}

func TestAppWithClient(t *testing.T) {
	a := New("ccl", "123456", WithName("代码工匠实验室"), WithClient("1", "2"))
	assert.Equal(t, "ccl", a.GetNamespaceName())
	assert.Equal(t, "123456", a.GetID())
	assert.Equal(t, "代码工匠实验室", a.GetName())
	assert.Equal(t, "1", a.GetClientID())
	assert.Equal(t, "2", a.GetClientSecret())
}

func TestAppOther(t *testing.T) {
	a := New("ccl", "123456", WithName("代码工匠实验室"),
		WithClient("1", "2"), WithAgent("000"), WithRobot("robot"))
	assert.Equal(t, "ccl", a.GetNamespaceName())
	assert.Equal(t, "123456", a.GetID())
	assert.Equal(t, "代码工匠实验室", a.GetName())
	assert.Equal(t, "1", a.GetClientID())
	assert.Equal(t, "2", a.GetClientSecret())
	assert.Equal(t, "000", a.GetAgentId())
	assert.Equal(t, "robot", a.GetRobotCode())
}

func TestGetAccessToken(t *testing.T) {
	if testClientSecret == "" || testClientId == "" {
		t.Skip("clientSecret or clientId is empty")
		return
	}
	ap := New("ccl", "118447d2-1c73-486f-8058-7daa046c9577",
		WithName("代码工匠实验室-监控平台"),
		WithClient(testClientId, testClientSecret),
		WithAgent("194334207"),
	)
	token := ap.GetAccessToken()
	assert.NotNil(t, token)
}

func TestGetAccessTokenWithRedis(t *testing.T) {
	if testClientSecret == "" || testClientId == "" {
		t.Skip("clientSecret or clientId is empty")
		return
	}
	ap := New("ccl", "118447d2-1c73-486f-8058-7daa046c9577",
		WithName("代码工匠实验室-监控平台"),
		WithClient(testClientId, testClientSecret),
		WithAgent("194334207"),
		WithRedis(redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "redis!@@&",
			DB:       0,
		}),
	)
	token := ap.GetAccessToken()
	assert.NotNil(t, token)
}
