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

func (User) TableName() string {
	return "user"
}

func GetUserInfoById(id int) map[string]interface{} {

	dbInit = getDb()

	user := User{}

	dbInit.First(&user, id)

	data := struct2Map(user)

	return data

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
