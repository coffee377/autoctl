package authority

import "math/big"

type OptType = uint

const (
	Add    OptType = 1
	Remove OptType = 2
)

func Opt(f Functions, points []FunctionPoint, opt OptType) Functions {
	bigInt := new(big.Int)
	for _, point := range points {
		if point == nil {
			continue
		}
		bigInt.Lsh(one, point.GetPosition()-1)
		switch opt {
		case Add:
			// 添加权限
			f.Get().Or(f.Get(), bigInt)
		case Remove:
			// 移除权限
			f.Get().AndNot(f.Get(), bigInt)
		}
	}
	return f
}
