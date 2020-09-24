package main

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"ginTest/model"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}

type Number1 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}

type Number2 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}

func main() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v", p)
			debug.PrintStack()
		}
	}()

	userId := 1
	orderInfoMap := map[string]interface{}{
		"GoodsId":  1,
		"MealTime": 111111111,
		"Total":    100,
		"Status":   0,
		"Remark":   "",
	}

	jsonStr := `[
				{"GoodsId":1,"Num":1,"goodsIngredientIds":"1,2,3"},
				{"GoodsId":2,"Num":12,"goodsIngredientIds":"1,3"}
	]`
	goodsInfoSlice := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonStr), &goodsInfoSlice)
	if err != nil {
		panic("json反序列化失败")
	}

	orderInfo := model.CreateOrder(userId, orderInfoMap)

	model.CreateOrderDetails(orderInfo.OrderNum, goodsInfoSlice)

	//
	// userInfo := map[string]interface{}{
	//	"nickname": "lkkame",
	// }

	// model.UpdateById(3, userInfo)
	// model.CreateUser(userInfo)

}
