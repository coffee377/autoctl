package authority

import (
	"errors"
)

var (
	PositionValueError = errors.New("功能点位置 position 的值必须为正整数")
)

type FunctionPoint interface {
	// GetName 功能点名称
	GetName() string
	// GetPosition 功能点位置，必须大于零，且不允许重复
	GetPosition() uint
	// GetDescription 功能点描述
	GetDescription() string
	//Authority
}
