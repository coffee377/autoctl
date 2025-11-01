package test

import (
	"cds/bid/data"
	"cds/bid/ds"
	"cds/dingtalk/app"
	"cds/dingtalk/oa"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpense(t *testing.T) {
	approval, err := oa.New(app.New("a57e9681-79cb-4242-96df-952be2dc3af7", app.WithRedis()))
	assert.Nil(t, err)

	client, ok := ds.Mysql()
	defer ds.CloseMysql(client)
	assert.Equal(t, true, ok)

	ctx := context.Background()
	id := "vSZoAivTSAub1LTfE3pVHA07201761911622"
	res, err := approval.GetProcessInstance(id)
	assert.Nil(t, err)

	url := fmt.Sprintf("https://applink.dingtalk.com/approval/detail?corpId=%s&instanceId=%s&from=%s",
		"dingd8b32bfb2b9da7b2", id, "dingopfniakkw72klkjv")
	t.Log(url)

	expenseData, err := data.NewBidExpense(id, res, data.WithUserHook(approval.GetUserHook()))
	assert.Nil(t, err)

	err = expenseData.Save(ctx, client)
	assert.Nil(t, err)
}

func TestExpenseBatch(t *testing.T) {
	approval, err := oa.New(app.New("a57e9681-79cb-4242-96df-952be2dc3af7", app.WithRedis()))
	assert.Nil(t, err)
	ids, err := approval.GetProcessInstanceIds(oa.BidExpenseProcessCode, "2025-01-01", "", nil)
	assert.Nil(t, err)

	client, ok := ds.Mysql()
	defer ds.CloseMysql(client)
	assert.Equal(t, true, ok)

	ctx := context.Background()

	for i, id := range ids {
		t.Logf("%d: %s", i+1, id)
		res, err := approval.GetProcessInstance(id)
		assert.Nil(t, err)

		expenseData, err := data.NewBidExpense(id, res, data.WithUserHook(approval.GetUserHook()))
		assert.Nil(t, err)

		err = expenseData.Save(ctx, client)
		assert.Nil(t, err)
	}
}
