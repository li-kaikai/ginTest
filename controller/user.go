package controller

import (
	"net/http"

	"ginTest/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetUserInfoById(context *gin.Context) {

	defer service.ReturnError(context)

	context.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"errMsg":  "success",
		"data":    service.GetUserInfoById(context),
	})

}

func CreateUser(context *gin.Context) {

	defer service.ReturnError(context)

	context.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"errMsg":  "success",
		"data":    service.Create(context),
	})

}

func UpdateById(context *gin.Context) {

	defer service.ReturnError(context)

	context.JSON(http.StatusOK, gin.H{
		"errCode": 0,
		"errMsg":  "success",
		"data":    service.UpdateById(context),
	})

}
