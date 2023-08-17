package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

type Action struct {
	Position    uint
	Name        string
	Description string
	Sort        int
}

func (a Action) GetSort() int {
	return a.Sort
}

func (a Action) Get() *big.Int {
	b := new(big.Int)
	lsh := b.Lsh(One, a.GetPosition()-1)
	return lsh
}

func (a Action) GetName() string {
	return a.Name
}

func (a Action) GetPosition() uint {
	return a.Position
}

func (a Action) GetDescription() string {
	return a.Description
}

var (
	QueryAction       = Action{1, "QUERY", "查询", 0}
	AddAction         = Action{2, "ADD", "增加", 0}
	DeleteAction      = Action{3, "DELETE", "删除", 0}
	BatchDeleteAction = Action{4, "BATCH_DELETE", "批量删除", 0}
	UpdateAction      = Action{5, "UPDATE", "更新", 0}
	ImportAction      = Action{6, "IMPORT", "导入", 0}
	ExportAction      = Action{7, "EXPORT", "导出", 0}
)

func TestWithFunctions(t *testing.T) {
	v := 0b00001111

	d := fmt.Sprintf("%d", v)
	authority := Authority(WithFunctions(d))

	result := authority.Has("0b00001100")
	assert.Equal(t, true, result)

	result = authority.Has("0b00010000")
	assert.Equal(t, false, result)

	result = authority.Has("14")
	assert.Equal(t, true, result)

	result = authority.HasAll(QueryAction, AddAction, DeleteAction, BatchDeleteAction)
	assert.Equal(t, true, result)

	result = authority.HasNone(UpdateAction, ImportAction, ExportAction)
	assert.Equal(t, true, result)

	result = authority.HasAny(UpdateAction, ImportAction, ExportAction, QueryAction, AddAction)
	assert.Equal(t, true, result)
}

func TestWithAuthority(t *testing.T) {
	authority := Authority(With(ExportAction))

	// 0b00000111
	authority.Add(QueryAction, AddAction, DeleteAction)

	assert.True(t, authority.Has("0b00000111"))
	assert.True(t, authority.Has("0b01000111"))

	authority.Remove(DeleteAction)
	authority.Remove(QueryAction)

	assert.False(t, authority.Has("0b00000111"))
	assert.True(t, authority.Has("0b01000000"))
	assert.True(t, authority.Has("0b00000010"))

	authority.Add(UpdateAction, ImportAction, ExportAction)

	assert.True(t, authority.Has(AddAction))
	assert.True(t, authority.Has(ExportAction))
	assert.True(t, authority.HasAll(ExportAction, AddAction))

	assert.False(t, authority.HasAll(QueryAction, AddAction, DeleteAction, BatchDeleteAction))
	assert.True(t, authority.HasAny(QueryAction, AddAction, DeleteAction, BatchDeleteAction))

	authority.Reset()
	assert.True(t, authority.Has(ExportAction))
	assert.True(t, authority.HasNone(QueryAction, AddAction, DeleteAction, BatchDeleteAction, UpdateAction, ImportAction))
}

func TestWithBinaryTextLength(t *testing.T) {
	authority := Authority(WithBinaryTextLength(16))
	// 0b1000011
	authority.Add(QueryAction, AddAction, ExportAction)
	d := authority.GetFunctions()
	assert.Equal(t, "67", d)
	b := authority.GetBaseFunctions(2)
	assert.Equal(t, "0000000001000011", b)
}
