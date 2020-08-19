package controller

import (
	"ginTest/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func GetUserInfoById(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": service.GetUserInfoById(context),
	})

}

func UpdateById(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": service.UpdateById(context),
	})

}