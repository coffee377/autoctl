package oa

import (
	"context"
	"fmt"
	"os"
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
	a := app.New("ccl", "118447d2-1c73-486f-8058-7daa046c9577",
		app.WithName("代码工匠实验室-监控平台"),
		app.WithClient(os.Getenv("APP_CLIENT_ID"), os.Getenv("APP_CLIENT_SECRET")),
		app.WithAgent("194334207"),
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
	ids, err2 := approval.GetProcessInstanceIds(context.Background(), BidApplyProcessCode, "2025-01-01", "", nil)
	assert.Nil(t, err2)
	assert.NotNil(t, ids)
	println(fmt.Sprintf("ids: %v", ids))
}

func TestGetProcessInstance(t *testing.T) {
	instance, err := approval.GetProcessInstance(context.TODO(), "NiXb3FWZRbKKyt4CFfQDmA07201743130227")
	assert.Nil(t, err)
	assert.NotNil(t, instance)
}
