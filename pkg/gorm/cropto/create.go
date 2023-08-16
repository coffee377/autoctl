package cropto

import "gorm.io/gorm"

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
