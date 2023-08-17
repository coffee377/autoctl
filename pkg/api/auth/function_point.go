package auth

import (
	"github.com/coffee377/autoctl/pkg/api/base"
	"github.com/coffee377/autoctl/pkg/api/dict"
)

type FunctionPoint interface {
	// IName 功能点名称
	dict.IName
	// GetPosition 功能点位置，必须大于零，且不允许重复
	GetPosition() uint
	// IDescription 功能点描述
	dict.IDescription
	base.ISort
}
