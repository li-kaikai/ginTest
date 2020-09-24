package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/mapstructure"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OrderDetail struct {
	Id                 int `gorm:"primary_key,column:id"`
	OrderNum           string
	GoodsId            int
	Num                int
	Total              int
	GoodsIngredientIds string
}

func (OrderDetail) TableName() string {
	return "order_details"
}

func CreateOrderDetails(orderNum string, orderDetails []map[string]interface{}) bool {

	for _, detail := range orderDetails {
		CreateOrderDetail(orderNum, detail)
	}

	return true

}

func CreateOrderDetail(orderNum string, detail map[string]interface{}) (orderDetail OrderDetail) {

	orderDetail = OrderDetail{OrderNum: orderNum}

	if err := mapstructure.Decode(detail, &orderDetail); err != nil {
		panic(err)
	}
	if re := db.Create(&orderDetail); re.Error != nil {
		panic(re.Error)
	}

	return

}
