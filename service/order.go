package service

import (
	"encoding/json"
	"strconv"

	"ginTest/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateOrder(context *gin.Context) bool {

	// userId
	userIdStr := context.Query("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// 订单信息
	orderInfoStr := context.Query("orderInfo")
	orderInfoByte := []byte(orderInfoStr)
	orderInfoMap := make(map[string]interface{})
	if err := json.Unmarshal(orderInfoByte, &orderInfoMap); err != nil {
		panic(err)
	}

	// 货品信息
	goodsInfoStr := context.Query("goodsInfo")
	goodsInfoSlice := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(goodsInfoStr), &goodsInfoSlice); err != nil {
		panic("json反序列化失败")
	}

	// 创建订单
	newOrderInfo := model.CreateOrder(userId, orderInfoMap)

	// 创建订单详情
	model.CreateOrderDetails(newOrderInfo.OrderNum, goodsInfoSlice)

	return true

}

func GetOrderInfoByOrderNum(ctx *gin.Context) (order map[string]interface{}) {

	orderNum := ctx.Param("orderNum")
	if orderNum == "" {
		orderNum = ctx.PostForm("orderNum")
		if orderNum == "" {
			panic("orderNum参数不能为空")
		}
	}

	order = model.GetInfoByOrderNum(orderNum)

	return
}

func UpdateOrderByOrderNum(ctx *gin.Context) {

}
