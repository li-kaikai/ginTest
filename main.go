package main

import (
    "ginTest/router"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

    router.Init()

}
