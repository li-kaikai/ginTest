package router

import (
	"ginTest/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() {

	router := gin.Default()

	router.Use(middleware.Logger())

	includeRouter(router)

	_ = router.Run(":8080")

}

func includeRouter(router *gin.Engine) {
	userRouter(router)
	authRouter(router)
}
