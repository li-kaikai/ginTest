package service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func ReturnError(context *gin.Context) {
	if p := recover(); p != nil {
		context.JSON(http.StatusOK, gin.H{
			"errNum": 1,
			"errMsg": "failed",
			"data":   false,
		})
	}
}
