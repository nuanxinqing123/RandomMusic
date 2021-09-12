package dataSource

import (
	"RandomMusic/model"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"os"
)

var DB *gorm.DB
var err error

func init() {
	mysqlConf := LoadMySQLConf()
	logoMode := mysqlConf.LogoMode
	//用户名：密码@tcp(ip:端口)/数据库?charset=utf8&parseTime&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime&loc=Local",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database,
	)

	DB, err = gorm.Open("mysql", dsn)

	//判断链接状态
	if err != nil {
		panic(err)
	}
	DB.LogMode(logoMode)

	//关闭空闲链接
	DB.DB().SetMaxOpenConns(100) //设置最大连接数
	DB.DB().SetMaxIdleConns(50)  //最大空闲数

	//自动迁移
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.UserEmailCode{})
	DB.AutoMigrate(&model.HistoryRecord{})
}

// LoadMySQLConf 读取配置文件
func LoadMySQLConf() *model.MySQLConf {
	mysqlConf := model.MySQLConf{}

	file, err := os.Open("conf/mysql.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(byteData, &mysqlConf)
	if err3 != nil {
		panic(err3)
	}

	return &mysqlConf
}
