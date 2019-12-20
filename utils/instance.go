package utils

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"os"
	"sync"
)

var (
	Config    *Configure   = nil
	DB        *xorm.Engine = nil
	mtxConfig sync.Mutex
	mtxDB     sync.Mutex
)

//加载配置文件，获取配置
func GetConfiguration(filename string) (*Configure, error) {
	mtxConfig.Lock()
	defer mtxConfig.Unlock()

	var conf Configure
	if Config == nil {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0666)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &conf)
		if err != nil {
			return nil, err
		}
		Config = &conf
	}
	return &conf, nil
}

func GetDB(dsn string) (*xorm.Engine, error) {
	mtxDB.Lock()
	defer mtxDB.Unlock()

	if DB == nil {
		//配置DB
		base, err := xorm.NewEngine("mysql", dsn)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		err = base.Ping()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		DB = base
	}
	return DB, nil
}
