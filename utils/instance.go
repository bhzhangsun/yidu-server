package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configure struct {
	Env   string `json:"env"`
	Debug string `json:"debug"`
	DB    string `json:"db"`
}

//加载配置文件，获取配置
func GetConfiguration(filename string) (*Configure, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var config Configure
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
