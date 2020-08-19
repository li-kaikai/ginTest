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
        userRouter.GET("getUserInfo/:id", controller.GetUserInfoById)
        userRouter.GET("updateUserInfo/:id/:name/:age", controller.UpdateById)
    }
}
