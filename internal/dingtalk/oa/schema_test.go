package oa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/buffer"
)

func TestGetBidApplyFormSchema(t *testing.T) {
	schema1, err1 := approval.GetFormSchema(BidApplyProcessCode)
	assert.Nil(t, err1)

	for i, item := range schema1.SchemaContent.Items {
		props := item.Props
		t.Log(i+1, *props.Label, *props.Id, *props.Required)
	}
}

func TestGetBidExpenseFormSchema(t *testing.T) {
	schema2, err2 := approval.GetFormSchema(BidExpenseProcessCode)
	assert.Nil(t, err2)

	buf := buffer.Buffer{}
	for i, item := range schema2.SchemaContent.Items {
		buf.Reset()
		props := item.Props
		pointer := false
		if props.Required != nil && *props.Required {
			pointer = true
		}
		_, _ = buf.WriteString(fmt.Sprintf("{ComponentId: \"%s\", FieldName: \"%s\", Converter: %s", *props.Id, *props.Label, "oa.StringConverter"))
		if pointer {
			_, _ = buf.WriteString(fmt.Sprintf(", Pointer: %t},", pointer))
		} else {
			_, _ = buf.WriteString("},")
		}
		t.Log(i+1, buf.String())
	}

}
