package oa

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	for i, item := range schema2.SchemaContent.Items {
		props := item.Props
		t.Log(i+1, *props.Label, *props.Id, *props.Required)
	}
}
