package dingtalk

import "context"

type Namespace interface {
	GetNamespaceName() string // 获取命名空间名称
}

type AccessToken interface {
	GetAccessToken(ctx context.Context) string
}
