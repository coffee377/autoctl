package dict

import "time"

type (
	ICode interface {
		// GetCode 编码
		GetCode() string
	}

	IName interface {
		// GetName 名称
		GetName() string
	}

	IValue[V any] interface {
		// GetValue 实际值
		GetValue() V
	}

	IText interface {
		// GetText 显示值
		GetText() string
	}

	IDescription interface {
		// GetDescription 描述
		GetDescription() string
	}
)

type Dictionary[Item DictionaryItem[any]] interface {
	ICode
	IName
	GetItems() []Item
}

type Dict struct {
	Id         uint      `json:"id,omitempty" gorm:"type:int;primaryKey"`
	CategoryId uint      `json:"cid,omitempty"`
	Code       string    `json:"code,omitempty"`
	Name       string    `json:"name,omitempty"`
	Sort       int       `json:"sort,omitempty"`
	Status     bool      `json:"status,omitempty"`
	Remark     string    `json:"remark,omitempty"`
	ExtJson    string    `json:"ext_json,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
