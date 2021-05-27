package service

import (
	"strconv"

	"ginTest/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetUserInfoById(context *gin.Context) (user map[string]interface{}) {

	id := context.Param("id")

	var idInt, err = strconv.Atoi(id)

	if err != nil {
		panic("error")
	}

	user = model.GetInfoById(idInt)

	return

}

func UpdateById(context *gin.Context) (user bool) {

	id := context.Param("id")
	nickname := context.Param("nickname")
	avatar := context.Param("avatar")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	userInfo := map[string]interface{}{
		"Nickname": nickname,
		"Avatar":   avatar,
	}

	user = model.UpdateById(idInt, userInfo)

	return

}

func Create(context *gin.Context) (user map[string]interface{}) {

	openid := context.Param("openid")
	nickname := context.Param("nickname")
	avatar := context.Param("avatar")
	gender := context.Param("gender")

	genderInt, err := strconv.Atoi(gender)
	if err != nil {
		panic(err.Error())
	}

	userInfo := map[string]interface{}{
		"openid":   openid,
		"nickname": nickname,
		"avatar":   avatar,
		"gender":   genderInt,
	}

	userId := model.CreateUser(userInfo)

	user = model.GetInfoById(userId)

	return

}
