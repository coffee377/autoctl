package auth

import (
	"math/big"
)

type Functions interface {

	// GetSource 原始功能点
	GetSource() *big.Int

	IAuthority

	// Reset 重置
	Reset()

	// GetFunctions 获取功能点字符串描述(10进制)
	// @return 功能字符串
	GetFunctions() string

	// GetBaseFunctions 获取功能点字符串描述
	// @param base 字符串表示形式的基数,必须在2到62之间
	// @return 功能字符串
	GetBaseFunctions(base int) string

	// Has 是否包含有某项功能
	// @param authority 功能
	// @return 是否包含指定功能
	Has(authority any) bool

	// HasAll 是否包含有所有功能
	// @param authority  功能
	// @param authorities 其他可选功能
	// @return 是否包含所有指定的功能点
	HasAll(authority IAuthority, authorities ...IAuthority) bool

	// HasAny 是否包含任何一项功能
	// @param authority  功能点
	// @param authorities 其他可选功能
	// @return 是否包含任何一项功能
	HasAny(authority IAuthority, authorities ...IAuthority) bool

	// HasNone 任何一项功能都没有
	// @param authority  功能
	// @param authorities 其他可选功能点
	// @return 是否包不包含所有的功能
	HasNone(authority IAuthority, authorities ...IAuthority) bool

	// Add 添加功能
	// @param authority  功能
	// @param authorities 其他可选功能
	// @return 功能点容器对象
	Add(authority IAuthority, authorities ...IAuthority) Functions

	// Remove 删除功能
	// @param authority  功能
	// @param authorities 其他可选功能
	// @return 功能点容器对象
	Remove(authority IAuthority, authorities ...IAuthority) Functions
}
