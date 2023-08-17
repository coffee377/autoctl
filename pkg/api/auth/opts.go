package auth

type OptType = uint

const (
	Add    OptType = 1
	Remove OptType = 2
)

func Opt(functions Functions, authorities []IAuthority, opt OptType) Functions {
	for _, authority := range authorities {
		if authority == nil {
			continue
		}
		switch opt {
		case Add:
			// 添加权限
			functions.Get().Or(functions.Get(), authority.Get())
		case Remove:
			// 移除权限
			functions.Get().AndNot(functions.Get(), authority.Get())
		}
	}
	return functions
}
