package service

import (
	"ginTest/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

func GetUserInfoById(context *gin.Context) (user map[string]interface{}) {

	id := context.Param("id")

	var idInt, err = strconv.Atoi(id)

	if err != nil {
		panic("error")
	}

	user = model.GetUserInfoById(idInt)

	//userJson, err2 := json.Marshal(user)
	//if err2 != nil {
	//	panic("user json error")
	//}
	//
	//userString := string(userJson)

	//m := make(map[string]interface{})
	//_ = json.Unmarshal(userJson, &m)

	return

}

func GetUserInfoTest(idInt int) map[string]interface{} {

	user := model.GetUserInfoById(idInt)

	return user
}
