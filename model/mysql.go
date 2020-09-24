package model

import (
	"ginTest/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {

	dbInit()

}

var db = dbInit()

func dbInit() (db *gorm.DB) {

	addr := config.Dft.Get().Mysql.Order.Addr

	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(err.Error())
	}

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	// 开启日志
	db.LogMode(true)

	return

}
