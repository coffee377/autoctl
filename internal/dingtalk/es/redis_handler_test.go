package es

import (
	"context"
	"testing"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestProduceMessage(t *testing.T) {
	header := event.EventHeader{
		EventCorpId: "dingd8b32bfb2b9da7b2",
		EventType:   "bpms_instance_change",
	}
	handler := redisEventHandler{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "redis!@@&",
			DB:       2,
		}),
	}
	rawData := `{
    "processInstanceId": "H7U9SGy0QAGUKu4_xClOcg07201768659795",
    "type": "terminate",
    "processCode": "PROC-958C3100-85BF-45D3-8583-6645DA922756"
  }`
	values := streamValues(header, []byte(rawData))
	err := handler.produceMessage(context.Background(), header, values)
	assert.Equal(t, nil, err)
}
