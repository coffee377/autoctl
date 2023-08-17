package dict

import (
	"github.com/coffee377/autoctl/pkg/api/base"
	"time"
)

type DictionaryCategory[ID string | uint] interface {
	base.ITreeSort[ID, DictionaryCategory[ID]]
	IName
	base.ICreatedAt
	base.IUpdatedAt
}

type Category struct {
	Id        uint       `json:"id,omitempty" gorm:"type:int;primaryKey"`
	Pid       uint       `json:"pid,omitempty"`
	Name      string     `json:"name,omitempty"`
	Sort      int        `json:"sort,omitempty"`
	Children  []Category `json:"children,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

func (c *Category) GetId() uint {
	return c.Id
}

func (c *Category) SetId(id uint) {
	c.Id = id
}

func (c *Category) GetParentId() uint {
	return c.Pid
}

func (c *Category) SetParentId(id uint) {
	c.Pid = id
}

func (c *Category) GetChildren() []Category {
	return c.Children
}

func (c *Category) GetSort() int {
	return c.Sort
}

func (c *Category) GetName() string {
	return c.Name
}

func (c *Category) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Category) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}
