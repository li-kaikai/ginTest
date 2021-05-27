package model

import (
	"encoding/json"
	"reflect"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/mapstructure"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Id       int `gorm:"primary_key"`
	Openid   string
	Nickname string
	Avatar   string
	Gender   int
}

func (User) TableName() string {
	return "user"
}

func GetInfoById(id int) map[string]interface{} {

	key := getUserRedisKey(id)

	redisRe := Redis.Get(key)
	if redisRe != false {
		userMap := redisRe2Map(redisRe)
		id := int(userMap["Id"].(float64))
		if id <= 0 {
			panic("用户不存在")
		}
	}

	user := User{}

	db.First(&user, id)

	if user.Id <= 0 {
		panic("用户不存在")
	}

	reflectT := reflect.TypeOf(user)
	reflectV := reflect.ValueOf(user)

	data := struct2Map(reflectT, reflectV)

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
	re := db.Model(&User{Id: id}).UpdateColumn(userInfo)
	if re.Error != nil {
		panic(re.Error)
	}

	key := getUserRedisKey(id)

	go func() {
		delRe := Redis.Del(key)
		redisChan <- delRe
	}()

	<-redisChan

	return true

}

func CreateUser(userInfo map[string]interface{}) int {

	user := User{}

	if err := mapstructure.Decode(userInfo, &user); err != nil {
		panic(err.Error())
	}

	re := db.Create(&user)

	if re.Error != nil {
		panic(re.Error)
	}

	id := 0
	if val, ok := re.Value.(*User); ok {
		id = val.Id
	}

	return id

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
