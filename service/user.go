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

    return

}

func UpdateById(context *gin.Context) (user bool) {

    id := context.Param("id")
    name := context.Param("name")
    age := context.Param("age")

    ageInt, err1 := strconv.Atoi(age)
    if err1 != nil {
        panic(err1.Error())
    }

    idInt, err2 := strconv.Atoi(id)
    if err2 != nil {
        panic(err1.Error())
    }

    userInfo := map[string]interface{}{
        "name": name,
        "age":  ageInt,
    }

    user = model.UpdateById(idInt, userInfo)

    return

}

func GetUserInfoTest(idInt int) map[string]interface{} {

    user := model.GetUserInfoById(idInt)

    return user
}
