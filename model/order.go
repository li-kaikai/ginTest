package model

import (
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/mapstructure"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Order struct {
	Id       int `gorm:"primary_key,column:id"`
	OrderNum string
	MealTime int
	Total    int
	Status   int
	UserId   int
	Remark   string
}

func (Order) TableName() string {
	return "orders"
}

// 创建订单
func CreateOrder(userId int, orderInfo map[string]interface{}) (order Order) {

	orderNum := GetOrderNum()

	order = Order{UserId: userId, OrderNum: orderNum}

	if err := mapstructure.Decode(orderInfo, &order); err != nil {
		panic(err)
	}

	if re := db.Create(&order); re.Error != nil {
		panic(re.Error)
	}

	return

}

// 获取订单号
func GetOrderNum() (orderNum string) {

	date := time.Now().Format("20060102150405")

	randNum := strconv.Itoa(getRandNum(1000, 9999))

	orderNum = date + randNum

	return

}

// 获取随机数
func getRandNum(min int, max int) int {

	return rand.Intn(max-min) + min

}
