package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//使用 gorm 定义model实体|结构体,没有采用 gorm.Model,生成的表是按下面的字段来创建的,即字段含 id、name、telephone、password
type User struct {
	ID        int    `grom:"AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:40;not null"`
}

func main() {

	db := InitDB()
	defer db.Close() // 延时关闭它

	r := gin.Default()
	r.POST("/api/auth/register", func(context *gin.Context) {
		//获取参数
		name := context.PostForm("name")
		telephone := context.PostForm("telephone")
		password := context.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 199,
				"msg":  "手机号必须是11位",
			})
			return
		}
		if len(password) < 6 {
			context.JSON(422, map[string]interface{}{
				"code": 199,
				"msg":  "密码必须大于6位",
			})
			return
		}
		//判断手机号是否已存在
		if isTelephoneExist(db, telephone) {
			context.JSON(422, gin.H{
				"code": 199,
				"msg":  "手机号已被占用,请换个试试",
			})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)

		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}

		//创建用户
		err := db.Create(&newUser).Error // 在插入对象 newUser 必须带 & 一同传入,否则会报错 using unaddressable value

		if err != nil {
			context.JSON(200, gin.H{
				"msg": "注册失败," + err.Error(),
			})
			return
		}

		//返回结果
		context.JSON(200, gin.H{
			"msg": "注册成功,手机号:" + telephone,
		})
	})
	panic(r.Run())
}

//查询手机号是否已存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	log.Println(user.ID) //输出 0
	db.Where("telephone = ?", telephone).First(&user)
	log.Println(user.ID)   //如果存在则输出 1,不存在打印 0
	log.Println(user.Name) //如果存在则输出对应的值 1,不存在打印 ""
	if user.ID != 0 {
		return true
	}
	return false
}

//随机生成10位的字符串
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnm123456987MNBVCXZASDFGHJKLPOIUYTREWQ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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
	db.AutoMigrate(&User{}) //自动创建表
	return db
}
