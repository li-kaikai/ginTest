package controller

import (
	"ginTest/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func GetUserInfoById(context *gin.Context) {

	defer service.ReturnError(context)

	context.JSON(http.StatusOK, gin.H{
		"errNum": 0,
		"errMsg": "success",
		"data":   service.GetUserInfoById(context),
	})

}

func CreateUser(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": service.Create(context),
	})

}

func UpdateById(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": service.UpdateById(context),
	})

}
