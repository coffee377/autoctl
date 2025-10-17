package oa

import (
	"testing"

	"github.com/coffee377/autoctl/internal/dingtalk/app"
	"github.com/coffee377/autoctl/pkg/log"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

const (
	BidApplyProcessCode   = "PROC-958C3100-85BF-45D3-8583-6645DA922756" // 投标申请审批表单编码
	BidExpenseProcessCode = "PROC-D8453B77-B313-4BEB-BE42-C71EE81DA61A" // 投标项目转款表单编码
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
	ids, err2 := approval.GetProcessInstanceIdsByMonth(BidApplyProcessCode, 2025, 1, nil)
	assert.Nil(t, err2)
	assert.NotNil(t, ids)
}

func TestGetProcessInstance(t *testing.T) {
	instance, err := approval.GetProcessInstance("NiXb3FWZRbKKyt4CFfQDmA07201743130227")
	assert.Nil(t, err)
	assert.NotNil(t, instance)
}
