package model

import (
	"ginTest/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {

	dbInit("GoTest")

}

func dbInit(dbName string) (db *gorm.DB) {

	var addr string
	if dbName == "GoTest2" {
		addr = config.Dft.Get().Mysql.GoTest2.Addr
	} else {
		addr = config.Dft.Get().Mysql.GoTest.Addr
	}

	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(err.Error())
	}

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	return

}
