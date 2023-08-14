package plugin

import (
	"github.com/coffee377/autoctl/pkg/log"
	"gorm.io/gorm"
	"reflect"
)

func callMethod(db *gorm.DB, fc func(value interface{}, tx *gorm.DB) bool) {
	tx := db.Session(&gorm.Session{NewDB: true})
	if called := fc(db.Statement.ReflectValue.Interface(), tx); !called {
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			db.Statement.CurDestIndex = 0
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				if value := reflect.Indirect(db.Statement.ReflectValue.Index(i)); value.CanAddr() {
					fc(value.Addr().Interface(), tx)
				} else {
					_ = db.AddError(gorm.ErrInvalidValue)
					return
				}
				db.Statement.CurDestIndex++
			}
		case reflect.Struct:
			if db.Statement.ReflectValue.CanAddr() {
				fc(db.Statement.ReflectValue.Addr().Interface(), tx)
			} else {
				_ = db.AddError(gorm.ErrInvalidValue)
			}
		}
	}
}

// BeforeCreatePassword before update password hooks
func BeforeCreatePassword(plugin CryptoPlugin) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		//if db.Error == nil && db.Statement.Schema != nil && !db.Statement.SkipHooks {
		//	fields := db.Statement.Schema.Fields
		//	for i, field := range fields {
		//		field.Set
		//	}
		//}

	}
}

// BeforeUpdatePassword before update password hooks
func BeforeUpdatePassword(db *gorm.DB) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if db.Error == nil && db.Statement.Schema != nil && !db.Statement.SkipHooks {
			callMethod(db, func(value interface{}, tx *gorm.DB) (called bool) {
				log.Warn("%v", value)

				//	db.AddError(i.BeforeUpdate(tx))
				return called
			})
		}

	}
}
