package cropto

import (
	"github.com/coffee377/autoctl/pkg/security/password"
	"gorm.io/gorm"
	"reflect"
)

const (
	// Name gorm plugin name
	Name = "crypto"
	// CryptoTag struct tag name
	CryptoTag = "crypto"
)

type CryptoPlugin struct {
	passwordEncoder password.Encoder
}

type Password interface {
	AfterQuery(*gorm.DB)
	BeforeCreate(*gorm.DB)
	BeforeUpdate(*gorm.DB)
}

func New() *CryptoPlugin {
	return &CryptoPlugin{
		password.CreateDelegatingPasswordEncoder(),
	}
}

func (c CryptoPlugin) Name() string {
	return Name
}

func (c CryptoPlugin) Initialize(db *gorm.DB) error {
	queryCallback := db.Callback().Query()
	_ = queryCallback.After("gorm:query").Register("crypt_plugin:before_create", c.AfterQuery)

	createCallback := db.Callback().Create()
	_ = createCallback.Before("gorm:create").Register("crypt_plugin:before_create", c.BeforeCreate)

	updateCallback := db.Callback().Update()
	_ = updateCallback.Before("gorm:update").Register("crypt_plugin:before_update", c.BeforeUpdate)
	return nil
}

func (c CryptoPlugin) AfterQuery(db *gorm.DB) {

}

func (c CryptoPlugin) BeforeCreate(db *gorm.DB) {
	BeforeCreatePassword(c)
}

func (c CryptoPlugin) BeforeUpdate(db *gorm.DB) {
	BeforeUpdatePassword(c)
}

type CryptoField struct {
	structField     reflect.StructField
	Type            string
	TypeStorageName string
	Salt            string
	SaltStorageName string
	Value           interface{}
}
