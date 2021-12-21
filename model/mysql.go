package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	DB, err = gorm.Open("mysql", "snowc:zkdAlert#@(111.62.122.250:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect msyql failed, err: %v \n", err)
		return
	}
	// 测试是否能够连通
	return DB.DB().Ping()
}
