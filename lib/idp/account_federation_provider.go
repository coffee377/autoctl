package idp

// AccountFederationProvider 联合身份提供商信息
type AccountFederationProvider struct {
	Id          uint   `json:"id,omitempty" gorm:"not null"`                       // 主键
	Code        uint8  `json:"code,omitempty" gorm:"type:int;not null"`            // 编码
	Name        string `json:"name,omitempty" gorm:"type:string;size:32;not null"` // 名称
	Description string `json:"description,omitempty" gorm:"type:string;size:128"`  // 描述
}
