package model

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"reflect"
	"strconv"
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

	key := getUserRedisKey(id)

	redisRe := Redis.Get(key)
	if redisRe != false {
		return redisRe2Map(redisRe)
	}

	user := User{}

	db.First(&user, id)

	data := struct2Map(user)

	redisChan := make(chan bool)

	go func() {
		setRe := setUserRedis(key, user)
		redisChan <- setRe
	}()

	<-redisChan

	return data

}

func UpdateById(id int, userInfo map[string]interface{}) bool {

	redisChan := make(chan bool)

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	db.Model(&User{Id: id}).UpdateColumn(userInfo)

	key := getUserRedisKey(id)

	go func() {
		delRe := Redis.Del(key)
		redisChan <- delRe
	}()

	<-redisChan

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

func redisRe2Map(redisRe interface{}) map[string]interface{} {
	dataByte := []byte(redisRe.(string))
	data := make(map[string]interface{})
	err := json.Unmarshal(dataByte, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func getUserRedisKey(id int) string {
	key := "user:" + strconv.Itoa(id)
	return key
}

func setUserRedis(key string, user User) bool {
	userJson, _ := json.Marshal(user)

	userString := string(userJson)

	return Redis.Set(key, userString, 0)
}
