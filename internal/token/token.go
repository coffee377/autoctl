package token

type Generator interface {
	// Generate 生成Token
	Generate(session Session) (string, error)
	// Validate 验证Token有效性
	Validate(token string) (bool, error)
}
