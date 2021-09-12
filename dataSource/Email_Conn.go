package dataSource

import (
	"RandomMusic/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

// LoadEmailConf 读取配置文件
func LoadEmailConf() *model.EmailConf {
	emialConf := model.EmailConf{}

	file, err := os.Open("conf/email.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(byteData, &emialConf)
	if err3 != nil {
		panic(err3)
	}

	return &emialConf
}
