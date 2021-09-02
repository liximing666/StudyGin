package main

import (
	"Gin/src/ORM/Tables"
	"Gin/tools"
)

func main() {

	//把连接数据库的方法写到工具类
	db := tools.ConnectDB("root", "root")
	defer db.Close()

	//初识CUDR

	//create or update
	user := Tables.User{Id: 4,Name: "fery", Age: 16, PassWord: "123"}
	if db.NewRecord(&user) == true {
		db.Create(&user)
	}else {
		db.Model(&user).Update("name", "cool")
	}



	db.First(&user)

	db.Model(&user).Update("name", "cool")


	db.Delete(&user)
	// 按照结构体属性生成表
	db.Table("info").CreateTable(&Tables.Info{})








}
