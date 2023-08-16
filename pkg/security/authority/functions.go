package authority

import (
	"github.com/todocoder/go-stream/stream"
	"math/big"
	"strings"
)

type Functions interface {
	// GetSource 原始功能点
	GetSource() *big.Int

	//// Count 功能点数量
	//Count() int

	// Reset 重置
	Reset()

	// GetFunctions 获取功能点字符串描述(10进制)
	// @return 功能字符串
	GetFunctions() string

	// GetBaseFunctions 获取功能点字符串描述
	// @param base 字符串表示形式的基数,必须在2到62之间
	// @return 功能字符串
	GetBaseFunctions(base int) string

	// Has 是否包含有某项功能点
	// @param point 功能点
	// @return 是否包含指定功能点
	Has(point FunctionPoint) bool

	// HasAll 是否包含有所有功能点
	// @param point  功能点
	// @param others 其他可选功能点
	// @return 是否包含所有指定的功能点
	HasAll(point FunctionPoint, others ...FunctionPoint) bool

	// HasAny 是否包含任何一项功能点
	// @param point  功能点
	// @param others 其他可选功能点
	// @return 是否包含任何一项功能点
	HasAny(point FunctionPoint, others ...FunctionPoint) bool

	// HasNone 任何一项功能点都没有
	// @param point  功能点
	// @param others 其他可选功能点
	// @return 是否包不包含所有的功能点
	HasNone(point FunctionPoint, others ...FunctionPoint) bool

	// Add 添加功能点
	// @param point  功能点
	// @param others 其他可选功能点
	// @return 功能点容器对象
	Add(point FunctionPoint, others ...FunctionPoint) Functions

	// Remove 删除功能点
	// @param point  功能点
	// @param others 其他可选功能点
	// @return 功能点容器对象
	Remove(point FunctionPoint, others ...FunctionPoint) Functions

	Authority
}

type DefaultFunctions struct {
	source *big.Int
	data   *big.Int
	//count  int
}

func (f *DefaultFunctions) GetSource() *big.Int {
	return f.source
}

//func (f *DefaultFunctions) Count() int {
//	return f.count
//}

func (f *DefaultFunctions) Reset() {
	f.data.SetString(f.source.String(), 10)
}

func (f *DefaultFunctions) GetFunctions() string {
	return f.GetBaseFunctions(10)
}

func (f *DefaultFunctions) GetBaseFunctions(base int) string {
	if base == 2 {
		text := f.data.Text(base)
		return strings.Repeat("0", 7-len(text)) + text
	}
	return f.data.Text(base)
}

func (f *DefaultFunctions) Has(point FunctionPoint) bool {
	if point != nil && zero.Cmp(f.data) == -1 {
		lsh := new(big.Int).Lsh(one, point.GetPosition()-1)
		r := lsh.And(lsh, f.data)
		cmp := r.Cmp(zero)
		return cmp == 1
	}
	return false
}

func (f *DefaultFunctions) HasAll(point FunctionPoint, others ...FunctionPoint) bool {
	return f.Has(point) && stream.Of(others...).AllMatch(func(fp FunctionPoint) bool {
		return f.Has(fp)
	})
}

func (f *DefaultFunctions) HasAny(point FunctionPoint, others ...FunctionPoint) bool {
	return f.Has(point) || stream.Of(others...).AnyMatch(func(fp FunctionPoint) bool {
		return f.Has(fp)
	})
}

func (f *DefaultFunctions) HasNone(point FunctionPoint, others ...FunctionPoint) bool {
	return !f.Has(point) && stream.Of(others...).NoneMatch(func(fp FunctionPoint) bool {
		return f.Has(fp)
	})
}

func (f *DefaultFunctions) Add(point FunctionPoint, others ...FunctionPoint) Functions {
	points := functionPoints(point, others)
	return Opt(f, points, Add)
}

func (f *DefaultFunctions) Remove(point FunctionPoint, others ...FunctionPoint) Functions {
	points := functionPoints(point, others)
	return Opt(f, points, Remove)
}

func (f *DefaultFunctions) Get() *big.Int {
	return f.data
}

func functionPoints(point FunctionPoint, others FunctionPoints) FunctionPoints {
	points := make(FunctionPoints, 0)
	if point != nil {
		points = append(points, point)
	}
	for _, other := range others {
		if other == nil {
			continue
		}
		points = append(points, other)
	}
	return points
}
