package router

import (
	"ginTest/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

func authRouter(router *gin.Engine) {
	userRouter := router.Group("auth")
	{
		userRouter.GET("getUserInfo/:id", func(context *gin.Context) {
			id := context.Param("id")
			var idInt, err = strconv.Atoi(id)
			if err != nil {
				panic("error")
			}

			user := model.GetUserInfoById(idInt)

			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": user,
			})
		})
	}
}
