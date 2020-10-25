package database

import (
	"fmt"

	"gorm.io/gorm/schema"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/conf"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var models = []interface{}{
	&model.User{},
}

// NewMySQL new db and retry connection when has error.
func InitMySQL(c *conf.Config) (db *gorm.DB, err error) {
	mysqlConfig := mysql.Config{
		DSN:                       c.Mysql.DSN, // DSN data source name
		DefaultStringSize:         191,         // string 类型字段的默认长度
		DisableDatetimePrecision:  true,        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,        // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,       // 根据版本自动配置
	}
	// gorm开启日志p配置
	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Error),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加复数形式，false默认加
		},
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		fmt.Println("gorm.Open failed, err:", err)
		return nil, err
	}

	if err = db.AutoMigrate(models...); nil != err {
	}
	return db, nil
}
