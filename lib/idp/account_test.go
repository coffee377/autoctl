package idp

import (
	"github.com/coffee377/autoctl/lib/idp/plugin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
)

func Test_MySql(d *testing.T) {
	p := 6
	config := mysql.Config{
		DSN: "root:root@1227@tcp(127.0.0.1:3306)/idass?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		//DefaultStringSize: 256,                                                                        // string 类型字段的默认长度
		//DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		DefaultDatetimePrecision:  &p,
	}
	mysqlDialect := mysql.New(config)
	db, err := gorm.Open(mysqlDialect, &gorm.Config{
		DryRun: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "idp_",
			SingularTable: true,
		},

		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.Use(plugin.NewCryptoPlugin())

	// 迁移 schema
	//err = db.AutoMigrate(&Account{}, &AccountFederation{}, &AccountFederationProvider{})
	if err != nil {
		panic("failed to auto migrate")
	}
	//db.Exec("truncate table accounts")
	var account Account

	affected := db.First(&account, "username = ?", "coffee377").RowsAffected

	if affected == 0 {
		// Create
		db.Create(&Account{Username: "coffee377"})
	}

	// Update - 将 account 的 password 更新为 123456
	//db.Debug().Model(&account).Where("username = ?", "coffee377").Update("password", "888888")
	// 更新密码
	db.Debug().Model(&account).Select("Password", "CryptoType", "CryptoSalt").Updates(&Account{Password: "test"})
}
