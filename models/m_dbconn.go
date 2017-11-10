package models

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"cloud9/config"
)

var Cloud10Db *sql.DB = nil

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UusConfig.Db.User, config.UusConfig.Db.Password, config.UusConfig.Db.Address,
		config.UusConfig.Db.Port, config.UusConfig.Db.DbName)
	Cloud10Db, _ = sql.Open("mysql", dsn)
}
