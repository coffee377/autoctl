package idp

import "time"

type AccountFederation struct {
	AccountId  uint      `json:"account_id,omitempty" gorm:"primaryKey;not null"`
	ProviderId uint      `json:"provider,omitempty" gorm:"primaryKey;not null"`
	ClientId   string    `json:"client_id,omitempty" gorm:"primaryKey;type:string;size:64;not null"`
	UnionId    string    `json:"union_id,omitempty" gorm:"primaryKey;type:string;size:64;not null"`
	OpenId     string    `json:"open_id,omitempty" gorm:"primaryKey;type:string;size:64;not null"`
	CreatedAt  time.Time `json:"created_at,omitempty"` // 创建时间
}
