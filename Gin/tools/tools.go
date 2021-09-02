package tools

import (
	"Gin/src/ORM/Tables"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func ToDate(timestamp int) string {

	timeobj := time.Unix(int64(timestamp), 0)

	return timeobj.Format("2006-01-02 15:04:05")
}

func GetDate() string {
	timeObj := time.Now()
	y := timeObj.Year()
	m := timeObj.Month()
	d := timeObj.Day()

	return fmt.Sprintf("%d_%s_%d", y, m, d)
}


func ConnectDB(root, password string) (*gorm.DB) {
	//连接db
	db, err := gorm.Open("mysql", root+":"+password+"@tcp(127.0.0.1:3306)/gromdb?charset=utf8&parseTime=True&loc=Local")
	//避免表名设置复数s
	db.SingularTable(true)

	if err != nil {
		panic(err)
	}

	//自动迁移创建的表
	db.AutoMigrate(&Tables.User{}, &Tables.CreditCard{})

	return db
}

