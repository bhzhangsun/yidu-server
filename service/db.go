package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"mobi.4se.tech/utils"
)

var (
	db *xorm.Engine
)

func init() {
	//配置文件读取
	config, err := utils.GetConfiguration("./config/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	//配置DB
	db, err := xorm.NewEngine("mysql", config.DB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}
