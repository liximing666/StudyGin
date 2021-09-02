package main

import (
	"Gin/src/ORM/Tables"
	"Gin/tools"
	"fmt"
)

func main() {

	db := tools.ConnectDB("root", "root")
	defer db.Close()


	//一般查询

	user := Tables.User{}
	db.First(&user)


	user1 := Tables.User{Id: 4}
	db.Find(&user1)

	users := []Tables.User{}
	db.Debug().Find(&users)
	fmt.Println(users)


	//条件查询
	user3 := Tables.User{}
	db.Debug().Where("name = ?", "aaa").Find(&user3)
	fmt.Println(user3)

	users1 := []Tables.User{}
	db.Where("name <> ?", "aaa").Find(&users1)//<> 是不等于
	fmt.Println(users1)

	//内联条件
	user5 := Tables.User{}
	db.Find(&user5, "name = ? and age = ?", "aaa", 16)
	fmt.Println(user5)


	//连接查询
	res := struct {
		Name string
		Age int
		CreditNum int
	}{}
	db.Model(&Tables.User{}).Joins("left join credit_card on user.id = credit_card.id").Where("name = ?", "aaa").Scan(&res)

	fmt.Println(res)













}
