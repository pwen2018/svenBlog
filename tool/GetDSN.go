package tool

import "fmt"

func GetDSN(cfg *Config) string {
	database := cfg.Database
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		database.User,
		database.Password,
		database.Host,
		database.Port,
		database.DbName,
		database.Charset,
		database.ParseTime,
		database.Loc,
	)
}
