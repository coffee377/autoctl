package idp

import (
	"time"
)

type Account struct {
	Id         uint      `json:"id,omitempty" gorm:"type:int;primaryKey"`
	Username   string    `json:"username,omitempty" gorm:"type:string;size:32;unique;not null;index:idx_account,priority:1"` // 用户名
	Password   string    `json:"password,omitempty" gorm:"type:string;size:128" crypto:"type:,crypto_type;salt:,"`           // 密码
	CryptoType string    `json:"crypto,omitempty"`
	CryptoSalt string    `json:"salt,omitempty"`
	RealName   string    `json:"real_name,omitempty" gorm:"type:string;size:8"`                                      // 真实姓名
	Mobile     string    `json:"mobile,omitempty" gorm:"type:string;size:11;index:idx_account,priority:2"`           // 手机号
	Email      string    `json:"email,omitempty" gorm:"type:string;size:128;index:idx_account,priority:3"`           // 邮箱
	Avatar     string    `json:"avatar,omitempty" gorm:"type:string;size:128"`                                       // 用户头像
	JobNumber  string    `json:"job_number,omitempty" gorm:"type:string;size:8;unique;index:idx_account,priority:4"` // 工号
	CreatedBy  uint      `json:"created_by,omitempty"`                                                               // 创建人
	CreatedAt  time.Time `json:"created_at,omitempty"`                                                               // 创建时间
	UpdatedBy  uint      `json:"updated_by,omitempty"`                                                               // 更新人
	UpdatedAt  time.Time `json:"updated_at,omitempty"`                                                               // 更新时间
	Flag       bool      `json:"flag,omitempty"`                                                                     // 逻辑删除标识
}

type Identity struct {
}
