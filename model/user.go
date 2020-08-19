package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"reflect"
)

type User struct {
	Id   int `gorm:"primary_key"`
	Name string
	Age  int
}

var db = dbInit("GoTest")

func (User) TableName() string {
	return "user"
}

func GetUserInfoById(id int) map[string]interface{} {

	user := User{}

	db.First(&user, id)

	data := struct2Map(user)

	return data

}

func UpdateById(id int, userInfo map[string]interface{}) bool {

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	db.Model(&User{Id: id}).UpdateColumn(userInfo)

	return true

}

func struct2Map(user User) (data map[string]interface{}) {

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	data = make(map[string]interface{})

	if t.NumField() > 0 {
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}

	return
}
