package token

import (
	"crypto/md5"
	"fmt"
	"time"

	"encoding/json"

	"github.com/google/uuid"
	"go.uber.org/zap/buffer"
)

type Session interface {
	GetId() string
	GetAccount() string
	GetStartUnixMilli() int64
	GetExpire() time.Duration
	json.Marshaler
	json.Unmarshaler
}

type jqSession struct {
	// @class: JSON中的@class字段，映射为Class
	Class string `json:"@class"`
	// 会话ID
	ID string `json:"id"`
	// 用户名
	Username string `json:"username"`
	// 认证请求类型
	AuthRequestType string `json:"authRequestType"`
	// 客户端IP
	IP *string `json:"ip"`
	// 会话开始时间：数组格式 [类型字符串, 时间戳]
	// 第一个元素："java.util.Date"，第二个元素：毫秒级时间戳（int64）
	StartTimestamp *int64 `json:"startTimestamp"`
	// 最后访问时间：毫秒级时间戳
	LastAccessTime *int64 `json:"lastAccessTime"`
	// 过期时间（秒）
	Expire time.Duration `json:"expire"`
	// 会话是否有效
	Valid bool `json:"valid"`
	// 失效原因（可为null，用*string表示可空）
	InvalidCause *string `json:"invalidCause"`
	// 会话附加属性：map类型，包含@class子字段
	Attrs map[string]interface{} `json:"attrs"`

	fixedId bool      // 是否固定会话 ID
	now     time.Time // 当前时间
}

func (j *jqSession) GetId() string {
	return j.ID
}

func (j *jqSession) GetAccount() string {
	return j.Username
}

func (j *jqSession) GetStartUnixMilli() int64 {
	return j.now.UnixMilli()
}

func (j *jqSession) GetExpire() time.Duration {
	return j.Expire
}

func (j *jqSession) MarshalJSON() ([]byte, error) {
	buf := buffer.Buffer{}
	_ = buf.WriteByte('{')
	_, _ = buf.WriteString(`"@class":"org.jsets.fastboot.security.session.Session"`)
	_, _ = buf.WriteString(fmt.Sprintf(`,"id":"%s"`, j.ID))
	_, _ = buf.WriteString(fmt.Sprintf(`,"username":"%s"`, j.Username))
	_, _ = buf.WriteString(fmt.Sprintf(`,"authRequestType":"%s"`, j.AuthRequestType))
	if j.IP != nil {
		_, _ = buf.WriteString(fmt.Sprintf(`,"ip":"%s"`, *j.IP))
	}
	if j.StartTimestamp != nil {
		_, _ = buf.WriteString(fmt.Sprintf(`,"startTimestamp":["java.util.Date",%d]`, *j.StartTimestamp))
	}
	if j.LastAccessTime != nil {
		_, _ = buf.WriteString(fmt.Sprintf(`,"lastAccessTime":%d`, *j.LastAccessTime))
	}
	_, _ = buf.WriteString(fmt.Sprintf(`,"expire":%d`, int(j.Expire.Seconds())))
	_, _ = buf.WriteString(fmt.Sprintf(`,"valid":%t`, j.Valid))
	if j.InvalidCause != nil {
		_, _ = buf.WriteString(fmt.Sprintf(`,"invalidCause":"%s"`, *j.InvalidCause))
	}
	attrs, _ := json.Marshal(j.Attrs)
	if len(attrs) > 0 {
		_, _ = buf.WriteString(`,"attrs":`)
		_, _ = buf.Write(attrs)
	}
	_ = buf.WriteByte('}')
	return buf.Bytes(), nil
}

func (j *jqSession) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

type SessionOptions func(*jqSession)

func NewJinQiSession(opts ...SessionOptions) Session {
	startTimestamp := time.Now().UnixMilli()
	session := &jqSession{
		Class:           "org.jsets.fastboot.security.session.Session",
		AuthRequestType: "UsernamePasswordRequest",
		StartTimestamp:  &startTimestamp,
		Expire:          time.Minute * 5,
		Valid:           true,
		Attrs: map[string]interface{}{
			"@class": "java.util.Collections$UnmodifiableMap",
		},
		now: time.Now(),
	}
	for _, opt := range opts {
		opt(session)
	}
	var sid string
	if session.fixedId {
		sum := md5.Sum([]byte(session.Username))
		sid = fmt.Sprintf("%x", sum)
	} else {
		v7, _ := uuid.NewV7()
		sid = v7.String()
	}
	session.ID = sid

	return session
}

func WithAccount(account string) SessionOptions {
	return func(s *jqSession) {
		s.Username = account
	}
}

func WithFixSession() SessionOptions {
	return func(s *jqSession) {
		s.fixedId = true
	}
}

func WithExpire(duration time.Duration) SessionOptions {
	return func(session *jqSession) {
		session.Expire = duration
	}
}
