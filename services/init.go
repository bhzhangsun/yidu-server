package services

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"yidu.4se.tech/config"
	"yidu.4se.tech/models"
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

	if err = base.Sync2(&models.Novel{}); err != nil {
		log.Println(err)
	}
	if err = base.Sync2(&models.Chapter{}); err != nil {
		log.Println(err)
	}
	DB = base
}
