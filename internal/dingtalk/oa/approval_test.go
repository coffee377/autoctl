package oa

import (
	"fmt"
	"testing"

	"github.com/coffee377/autoctl/internal/dingtalk"
	"github.com/coffee377/autoctl/internal/dingtalk/utils"
	"github.com/stretchr/testify/assert"
)

var (
	approval *Approval
	err      error
)

func init() {
	approval, err = New(dingtalk.App{
		Id:           "a57e9681-79cb-4242-96df-952be2dc3af7",
		Name:         utils.ToPtr("安徽晶奇-统一认证"),
		AgentId:      "1038540627",
		ClientID:     "dingopfniakkw72klkjv",
		ClientSecret: "6Il0DuPZPPIr-OG03uMrnqDNu_o03tpIkK03ScpuEPP6NAw7J52D0LWPvTjRf4BR",
	})
	if err != nil {
		panic(err)
	}
}

func TestGetProcessInstanceIds(t *testing.T) {
	ids, err2 := approval.GetProcessInstanceIds(BidApplyProcessCode, "2025-01-01", "", nil)
	assert.Nil(t, err2)
	assert.NotNil(t, ids)
	println(fmt.Sprintf("ids: %v", ids))
}

func TestGetProcessInstance(t *testing.T) {
	instance, err := approval.GetProcessInstance("NiXb3FWZRbKKyt4CFfQDmA07201743130227")
	assert.Nil(t, err)
	assert.NotNil(t, instance)
}
