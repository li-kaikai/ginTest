package model

import (
	"ginTest/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbInit *gorm.DB

func init() {

	dbInit = getDb()

	_ = dbInit.Close()

}

func getDb() *gorm.DB {

	dbInit, err := gorm.Open("mysql", config.Dft.Get().Mysql.GoTest.Addr)
	if err != nil {
		panic(err.Error())
	}

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	dbInit.SingularTable(true)

	return dbInit

}
