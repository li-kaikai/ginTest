package main

import (
	"encoding/json"
	"fmt"
	"ginTest/service"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	user := service.GetUserInfoTest(1)

	//println("%T", user)

	userJson, err := json.Marshal(user)
	if err != nil {
		panic("user json error")
	}

	a := string(userJson)

	fmt.Println(a)

	//fmt.Printf("%+v\n", userJson)
	fmt.Printf("%T\n", a)

}
