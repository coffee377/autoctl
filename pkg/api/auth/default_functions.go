package auth

import (
	"fmt"
	"github.com/coffee377/autoctl/pkg/log"
	"github.com/todocoder/go-stream/stream"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type defaultFunctions struct {
	source        *big.Int
	data          *big.Int
	binaryTextLen int
}

func (f *defaultFunctions) GetSource() *big.Int {
	return f.source
}

func (f *defaultFunctions) Reset() {
	f.data.SetString(f.source.String(), 10)
}

func (f *defaultFunctions) GetFunctions() string {
	return f.GetBaseFunctions(10)
}

func (f *defaultFunctions) GetBaseFunctions(base int) string {
	if base == 2 {
		text := f.data.Text(base)
		l := len(text)
		if f.binaryTextLen > 0 {
			l = f.binaryTextLen
		}
		format := strings.Join([]string{"%0", strconv.Itoa(l), "s"}, "")
		return fmt.Sprintf(format, text)
	}
	return f.data.Text(base)
}

func (f *defaultFunctions) Has(authority any) bool {
	if authority != nil && Zero.Cmp(f.data) == -1 {
		var auth *big.Int
		switch res := authority.(type) {
		case IAuthority:
			auth = res.Get()
			break
		case FunctionPoint:
			auth = new(big.Int).Lsh(One, res.GetPosition()-1)
		case string:
			auth, _ = new(big.Int).SetString(res, 0)
			break
		default:

		}

		if log.IsDebugEnabled() {
			t1 := f.data.Text(2)
			t2 := auth.Text(2)
			l := math.Max(float64(len(t1)), float64(len(t2)))
			format := strings.Join([]string{"%0", strconv.Itoa(int(l)), "s"}, "")
			log.Debug("%s & %s", fmt.Sprintf(format, t1), fmt.Sprintf(format, t2))
		}

		return has(f.data, auth)
	}
	return false
}

func has(data *big.Int, auth *big.Int) bool {
	return data != nil && auth != nil && new(big.Int).And(auth, data).Cmp(auth) == 0
}

func (f *defaultFunctions) HasAll(authority IAuthority, authorities ...IAuthority) bool {
	return f.Has(authority) && stream.Of(authorities...).AllMatch(func(fp IAuthority) bool {
		return f.Has(fp)
	})
}

func (f *defaultFunctions) HasAny(authority IAuthority, authorities ...IAuthority) bool {
	return f.Has(authority) || stream.Of(authorities...).AnyMatch(func(fp IAuthority) bool {
		return f.Has(fp)
	})
}

func (f *defaultFunctions) HasNone(authority IAuthority, authorities ...IAuthority) bool {
	return !f.Has(authority) && stream.Of(authorities...).NoneMatch(func(fp IAuthority) bool {
		return f.Has(fp)
	})
}

func (f *defaultFunctions) Add(authority IAuthority, authorities ...IAuthority) Functions {
	auths := mergeAuthorities(authority, authorities)
	return Opt(f, auths, Add)
}

func (f *defaultFunctions) Remove(authority IAuthority, authorities ...IAuthority) Functions {
	auths := mergeAuthorities(authority, authorities)
	return Opt(f, auths, Remove)
}

func (f *defaultFunctions) Get() *big.Int {
	return f.data
}

func mergeAuthorities(point IAuthority, others []IAuthority) []IAuthority {
	authorities := make([]IAuthority, 0)
	if point != nil {
		authorities = append(authorities, point)
	}
	for _, other := range others {
		if other == nil {
			continue
		}
		authorities = append(authorities, other)
	}
	return authorities
}
