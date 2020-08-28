package database

import (
	"com.fwtai/app2/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasouce.driverName")
	host := viper.GetString("datasouce.host")
	port := viper.GetString("datasouce.port")
	database := viper.GetString("datasouce.database")
	username := viper.GetString("datasouce.username")
	password := viper.GetString("datasouce.password")
	charset := viper.GetString("datasouce.charset")
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
