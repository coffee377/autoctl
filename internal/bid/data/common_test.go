package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtraDictCode(t *testing.T) {
	workflowData := NewWorkflowData("", nil)
	typ1 := "代理服务费"
	code1, ok1 := workflowData.ExtraDictCode(&typ1)
	assert.Equal(t, false, ok1)
	assert.NotNil(t, "代理服务费", code1)

	typ2 := "标书工本费(DF)"
	code2, ok2 := workflowData.ExtraDictCode(&typ2)
	assert.Equal(t, true, ok2)
	assert.NotNil(t, "DF", code2)

}
