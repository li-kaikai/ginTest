package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ReturnError(context *gin.Context) {
	if p := recover(); p != nil {
		context.JSON(http.StatusOK, gin.H{
			"errCode": 1,
			"errMsg":  p,
			"data":    false,
		})
	}
}
