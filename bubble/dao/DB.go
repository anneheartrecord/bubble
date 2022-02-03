package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golangstudy/gowebjinjie/bubble/model"
)
var DB *gorm.DB
func InitDB() (err error) {
	dsn:="root:cxs20030416@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true"
	db,err:=gorm.Open("mysql",dsn)
	if err!=nil {
		fmt.Println("failed to connect to the db")
		return
	}
	DB=db
	DB.AutoMigrate(&model.Todo{})
	return
}