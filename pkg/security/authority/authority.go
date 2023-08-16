package authority

import (
	"github.com/todocoder/go-stream/stream"
	"math/big"
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

type Authority interface {
	Get() *big.Int
}

func WithAuthority(authority ...FunctionPoint) Functions {
	bigInt := new(big.Int)

	source := new(big.Int)
	data := new(big.Int)

	stream.Of(authority...).Filter(func(item FunctionPoint) bool {
		return item != nil
	}).ForEach(func(item FunctionPoint) {
		bigInt.Lsh(one, item.GetPosition()-1)
		data.Or(data, bigInt)
	})

	source.SetString(data.String(), 10)

	functions := &DefaultFunctions{source, data}
	return functions
}

func WithFunctions(functions string, base int) Functions {
	source := new(big.Int)
	data := new(big.Int)
	source.SetString(functions, base)
	data.SetString(functions, base)
	fns := DefaultFunctions{source, data}
	return &fns
}
