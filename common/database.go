package common

import (
	"com.fwtai/app2/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "192.168.3.66"
	port := "3306"
	database := "golang"
	username := "root"
	password := "rootFwtai"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接数据库失败,原因:" + err.Error())
	}
	db.AutoMigrate(&model.User{}) //自动创建表
	DB = db
	return db
}

//返回实例
func GetDB() *gorm.DB {
	return DB
}
