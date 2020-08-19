package main

import (
	"ginTest/model"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}

type Number1 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}

type Number2 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}

func main() {

	userInfo := model.GetUserInfoById(1)

	println(userInfo)

	//age := "18"
	//id := "2"
	//name := "lkklkk"
	//
	//ageInt, err1 := strconv.Atoi(age)
	//if err1 != nil {
	//    panic(err1.Error())
	//}
	//
	//idInt, err2 := strconv.Atoi(id)
	//if err2 != nil {
	//    panic(err1.Error())
	//}
	//
	//userInfo := map[string]interface{}{
	//    "name": name,
	//    "age":  ageInt,
	//}
	//
	//_ = model.UpdateById(idInt, userInfo)
	//
	//println(idInt)

	//user := service.GetUserInfoTest(1)
	//
	////println("%T", user)
	//
	//userJson, err := json.Marshal(user)
	//if err != nil {
	//    panic("user json error")
	//}
	//
	//a := string(userJson)
	//
	//fmt.Println(a)
	//
	////fmt.Printf("%+v\n", userJson)
	//fmt.Printf("%T\n", a)

}
