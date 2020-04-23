package services

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"yidu.4se.tech/config"
)

var DB *xorm.Engine = nil

func init() {
	//配置DB
	base, err := xorm.NewEngine("mysql", config.MYSQL)
	if err != nil {
		fmt.Println("NewEngine:", err)
	}

	err = base.Ping()
	if err != nil {
		fmt.Println("Ping:", err)
	}
	DB = base
}
