package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Client *gorm.DB

func init()  {
	InitDB()
}

func InitDB()  {
	var err error
	Client, err = gorm.Open("mysql", "root:kangkang@(127.0.0.1:3306)/gohello?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
