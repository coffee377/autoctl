package oa

import (
	"cds/dingtalk/app"
	"fmt"
	"testing"

	"github.com/coffee377/autoctl/pkg/log"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/buffer"
)

var (
	approval *Approval
	err      error
)

func init() {
	a := app.New("a57e9681-79cb-4242-96df-952be2dc3af7",
		app.WithRedis(redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "redis!@@&",
			DB:       0,
		}),
	)
	approval, err = New(a)
	if err != nil {
		log.Error(err.Error())
	}
}

func TestGetProcessInstanceIds(t *testing.T) {
	ids, err2 := approval.GetProcessInstanceIds(BidApplyProcessCode, "2025-01-01", "", nil)
	assert.Nil(t, err2)
	assert.NotNil(t, ids)
	buf := buffer.Buffer{}
	t.Logf("共 %d 条记录", len(ids))
	_, _ = buf.WriteString("\ninsert into ids (id) values ")
	for i, id := range ids {
		_, _ = buf.WriteString(fmt.Sprintf("('%s')", id))
		if i < len(ids)-1 {
			_ = buf.WriteByte(',')
		} else {
			_ = buf.WriteByte(';')
		}
	}
	t.Logf("%s", buf.String())
}

func TestGetProcessInstance(t *testing.T) {
	instId := "Rp9D_t0WQrqgpfxvxUZ_EQ07201760318742"
	res, err := approval.GetProcessInstance(instId)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	status := res.Status
	assert.Equal(t, "COMPLETED", *status)
	result := res.Result
	assert.Equal(t, "agree", *result)
}
