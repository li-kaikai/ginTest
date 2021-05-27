package model

import (
	"encoding/json"
	"reflect"

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

func redisRe2Map(redisRe interface{}) map[string]interface{} {
	dataByte := []byte(redisRe.(string))
	data := make(map[string]interface{})
	err := json.Unmarshal(dataByte, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func struct2Map(t reflect.Type, v reflect.Value) (data map[string]interface{}) {

	data = make(map[string]interface{})

	if t.NumField() > 0 {
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}

	return
}
