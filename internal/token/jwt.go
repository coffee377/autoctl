package token

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// JWT类型的Token生成器
type jwtGenerator struct {
	Generator
	issuer         string            // 签发者
	signMethod     jwt.SigningMethod // 签名方法
	signKey        []byte            // 签名密钥
	audience       []string          // 受众
	b64            bool              // 是否对 token base64编码
	tokenKeyPrefix string
	rdb            *redis.Client
}

// NewJWTGenerator 创建JWT生成器实例
func NewJWTGenerator(opts ...JWTGeneratorOptions) Generator {
	j := &jwtGenerator{
		issuer:         "ccl.site",
		signMethod:     jwt.SigningMethodHS256,
		signKey:        []byte("qwertyuiop1234567890"),
		audience:       []string{"ccl"},
		b64:            false,
		tokenKeyPrefix: "token:",
		rdb: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "redis!@@&",
		}),
	}
	for _, opt := range opts {
		opt(j)
	}
	return j
}

// CustomClaims 自定义Claims（包含标准注册Claims + 自定义字段）
type CustomClaims struct {
	// 自定义字段（根据业务需求添加）
	Username string `json:"username"`
	// 嵌入标准注册Claims（必须，包含exp/iss/aud等核心验证字段）
	jwt.RegisteredClaims
}

// Generate 实现 TokenGenerator 接口的 Generate 方法
func (j *jwtGenerator) Generate(session Session) (string, error) {
	// 构造 Claims
	v7, _ := uuid.NewV7()
	now := time.UnixMilli(session.GetStartUnixMilli())
	claims := CustomClaims{
		Username: session.GetAccount(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   session.GetId(),
			NotBefore: jwt.NewNumericDate(now), // 生效时间（立即生效）
			IssuedAt:  jwt.NewNumericDate(now), // 签发时间
			ID:        v7.String(),
		},
	}
	if j.issuer != "" {
		claims.Issuer = j.issuer
	}
	if len(j.audience) > 0 {
		claims.Audience = j.audience
	}
	expire := session.GetExpire()
	if expire > 0 {
		claims.ExpiresAt = jwt.NewNumericDate(now.Add(expire))
	}

	// 签发JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.signKey)
	if err != nil {
		return "", err
	}

	if j.b64 {
		signedToken = base64.URLEncoding.EncodeToString([]byte(signedToken))
	}

	sessionBytes, _ := json.Marshal(session)
	_, err = j.rdb.Set(context.TODO(), fmt.Sprintf("%s%s", j.tokenKeyPrefix, session.GetId()), string(sessionBytes), session.GetExpire()).Result()
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Validate 实现 TokenGenerator 接口的 Validate 方法
func (j *jwtGenerator) Validate(tokenStr string) (bool, error) {
	return false, nil
}

type JWTGeneratorOptions func(*jwtGenerator)

func WithIssuer(issuer string) JWTGeneratorOptions {
	return func(j *jwtGenerator) {
		j.issuer = issuer
	}
}

func WithSign(m jwt.SigningMethod, key []byte) JWTGeneratorOptions {
	return func(j *jwtGenerator) {
		j.signKey = key
		j.signMethod = m
	}
}

func WithAudience(audience ...string) JWTGeneratorOptions {
	return func(generator *jwtGenerator) {
		generator.audience = audience
	}
}

func WithTokenBase64Encode() JWTGeneratorOptions {
	return func(generator *jwtGenerator) {
		generator.b64 = true
	}
}

func WithTokenPrefix(prefix string) JWTGeneratorOptions {
	return func(generator *jwtGenerator) {
		generator.tokenKeyPrefix = prefix
	}
}

func WithRedis(options redis.Options) JWTGeneratorOptions {
	return func(generator *jwtGenerator) {
		generator.rdb = redis.NewClient(&options)
	}
}

func WithJinQi() JWTGeneratorOptions {
	return func(generator *jwtGenerator) {
		WithIssuer("JinQi")(generator)
		WithSign(jwt.SigningMethodHS256, []byte("abcdee78985eve*86sfdafec"))(generator)
		WithAudience("cds")(generator)
		WithTokenBase64Encode()(generator)
		WithTokenPrefix("fastboot-session::")(generator)
	}
}
