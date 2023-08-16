package cropto

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func parsingCryptoField(structField reflect.StructField, value interface{}) (*CryptoField, bool) {
	tag := structField.Tag.Get(CryptoTag)
	cryptoField := &CryptoField{
		structField:     structField,
		Type:            "bcrypt",
		TypeStorageName: "CryptoType",
		Salt:            "",
		SaltStorageName: "CryptoSalt",
		Value:           value,
	}
	if tag != "" {
		split := strings.Split(tag, ";")
		for _, s := range split {
			if strings.TrimSpace(s) == "" {
				continue
			}
			i := strings.Split(s, ":")
			v := strings.TrimSpace(i[1])
			if v == "" {
				continue
			}
			k := strings.TrimSpace(i[0])
			switch k {
			case "type":
				i := strings.Index(v, ",")
				if i == -1 {
					if v != "" {
						cryptoField.Type = v
					}

				} else {
					if v[:i] != "" {
						cryptoField.Type = v[:i]
					}
					if v[i+1:] != "" {
						cryptoField.TypeStorageName = v[i+1:]
					}
				}
				break
			case "salt":
				i := strings.Index(v, ",")
				if i == -1 {
					if v != "" {
						cryptoField.Salt = v
					}
				} else {
					if v[:i] != "" {
						cryptoField.Salt = v[:i]
					}
					if v[i+1:] != "" {
						cryptoField.SaltStorageName = v[i+1:]
					}
				}
				break
			}
		}
		return cryptoField, true
	} else if structField.Name == "Password" {
		// 对于有 Password,ClientSecret 字段的模型，更新记录时，将该字段的值按照默认密码方式进行加密处理
		return cryptoField, true
	}
	return nil, false
}

func getReflectElem(i interface{}) (reflect.Type, reflect.Value) {
	destType := reflect.TypeOf(i)
	destValue := reflect.ValueOf(i)

	for destType.Kind() == reflect.Pointer {
		destType = destType.Elem()
		destValue = destValue.Elem()
	}

	return destType, destValue
}

// BeforeUpdatePassword before update password hooks
func BeforeUpdatePassword(plugin CryptoPlugin) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		dbSchema := db.Statement.Schema
		if db.Error == nil && dbSchema != nil {
			dest := db.Statement.Dest
			cryptoFieldMap := make(map[string]*CryptoField)
			// 更新数据类型是 map
			if updateInfo, ok := dest.(map[string]interface{}); ok {
				for columnName, columnValue := range updateInfo {
					field := dbSchema.LookUpField(columnName)
					if field == nil {
						continue
					}
					if cryptoField, ok := parsingCryptoField(field.StructField, columnValue); ok {
						cryptoFieldMap[field.DBName] = cryptoField
					}
				}
				for columnName, cryptoField := range cryptoFieldMap {
					encoded := plugin.passwordEncoder.Encode(cast.ToString(cryptoField.Value))
					updateInfo[columnName] = encoded
				}
				return
			}

			destType, destValue := getReflectElem(db.Statement.Dest)
			if destType != nil {
				for i := 0; i < destType.NumField(); i++ {
					structField := destType.Field(i)
					val := destValue.Field(i)
					if len(val.String()) == 0 {
						continue
					}
					if cryptoField, ok := parsingCryptoField(structField, val); ok {
						//cryptoFieldMap[structField.Name] = cryptoField
						// 1. 更新密码
						if val.CanSet() {
							// 对密码进行加密处理
							encoded := plugin.passwordEncoder.Encode(cast.ToString(cryptoField.Value))
							val.SetString(encoded)
							//val.Set(reflect.ValueOf(encoded))
						}

						// 2. 更新加密算法
						cryptoType := dbSchema.LookUpField(cryptoField.TypeStorageName)
						if cryptoType != nil {
							value := destValue.FieldByName(cryptoType.StructField.Name)
							//_ = cryptoType.Set(db.Statement.Context, value, "newValue")
							//field := cryptoType.StructField
							if value.CanSet() {
								value.Set(reflect.ValueOf(cryptoField.Type))
							}
						}

						// 3. 更新 salt
						cryptoSalt := dbSchema.LookUpField(cryptoField.SaltStorageName)
						if cryptoSalt != nil {
							value := destValue.FieldByName(cryptoSalt.StructField.Name)
							if value.CanSet() {
								value.Set(reflect.ValueOf(cryptoField.Salt))
							}
						}
					}
				}
			}
		}
	}
}
