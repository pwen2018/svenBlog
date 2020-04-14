package tool

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitGorm(cfg *Config, dsn string) (err error) {
	database := cfg.Database
	DB, _ = gorm.Open(database.Drive, dsn)
	if database.Debug {
		DB = DB.Debug()
	}
	err = DB.DB().Ping()
	if err != nil {
		return err
	}
	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	DB.SingularTable(database.Singular)
	return err
}
