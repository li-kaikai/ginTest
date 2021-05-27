package router

import (
	"ginTest/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func userRouter(router *gin.Engine) {

	userRouter := router.Group("user")
	{
		userRouter.GET("get/:id", controller.GetUserInfoById)
		userRouter.GET("update/:id/:nickname/:avatar", controller.UpdateById)
		userRouter.GET("create/:openid/:nickname/:avatar/:gender", controller.CreateUser)
	}

}
