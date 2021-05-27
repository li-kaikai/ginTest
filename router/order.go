package router

import (
	"ginTest/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func orderRouter(router *gin.Engine) {

	userRouter := router.Group("order")
	{
		userRouter.GET("get/:orderNum", controller.GetOrderInfoByOrderNum)
		userRouter.POST("get", controller.GetOrderInfoByOrderNum)
		userRouter.GET("create", controller.CreateOrder)
	}

}
