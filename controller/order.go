package controller

import (
	"net/http"

	"ginTest/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetOrderInfoByOrderNum(context *gin.Context) {

	defer service.ReturnError(context)

	context.JSON(http.StatusOK, gin.H{
		"errNum": 0,
		"errMsg": "success",
		"data":   service.GetOrderInfoByOrderNum,
	})

}

func CreateOrder(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": service.CreateOrder(context),
	})

}

func UpdateOrderByOrderNum(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": service.UpdateOrderByOrderNum,
	})

}
