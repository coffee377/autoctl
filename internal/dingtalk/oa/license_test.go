package oa

import (
	"testing"
	"time"

	"github.com/coffee377/autoctl/internal/dingtalk/app"
	"github.com/stretchr/testify/assert"
)

func Test_Lic(t *testing.T) {
	oa, err := NewLic(app.New("a57e9681-79cb-4242-96df-952be2dc3af7",
		app.WithRedis(),
	))
	assert.Nil(t, err)
	startTime := time.Now().Format(time.DateOnly)
	//startTime = "2025-01-01"
	err = oa.Run(startTime, "")
	assert.Nil(t, err, "err:%v")
}
