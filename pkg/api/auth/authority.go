package auth

import (
	"github.com/todocoder/go-stream/stream"
	"math/big"
)

var (
	Zero = big.NewInt(0)
	One  = big.NewInt(1)
)

type IAuthority interface {
	Get() *big.Int
}

type AuthorityOption func(functions *defaultFunctions)

func Authority(opts ...AuthorityOption) Functions {
	fns := &defaultFunctions{source: new(big.Int), data: new(big.Int), binaryTextLen: 0}
	for _, opt := range opts {
		opt(fns)
	}
	return fns
}

func With(authority ...IAuthority) AuthorityOption {
	data := new(big.Int)

	stream.Of(authority...).Filter(func(item IAuthority) bool {
		return item != nil
	}).ForEach(func(item IAuthority) {
		data.Or(data, item.Get())
	})

	source := new(big.Int)
	source.SetString(data.String(), 10)

	return func(functions *defaultFunctions) {
		functions.source = source
		functions.data = data
	}
}

func WithFunctions(functions string) AuthorityOption {
	source := new(big.Int)
	data := new(big.Int)
	source.SetString(functions, 0)
	data.SetString(functions, 0)
	return func(functions *defaultFunctions) {
		functions.source = source
		functions.data = data
	}
}

func WithBinaryTextLength(len int) AuthorityOption {
	return func(functions *defaultFunctions) {
		functions.binaryTextLen = len
	}
}
