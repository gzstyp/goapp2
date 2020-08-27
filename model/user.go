package model

import "github.com/jinzhu/gorm"

//使用 gorm 定义model实体|结构体
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:40;not null"`
}
