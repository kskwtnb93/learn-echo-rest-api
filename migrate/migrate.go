package main

import (
	"fmt"
	"learn-echo-rest-api/db"
	"learn-echo-rest-api/model"
)

func main() {
	dbConn := db.NewDB() // データベースのアドレス
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
