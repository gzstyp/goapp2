package controller

import (
	"com.fwtai/app2/common"
	"com.fwtai/app2/model"
	"com.fwtai/app2/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	DB := common.GetDB() //前面的 DB 写成db小写也是可以的,下面要跟着改
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
	if isTelephoneExist(DB, telephone) {
		context.JSON(422, gin.H{
			"code": 199,
			"msg":  "手机号已被占用,请换个试试",
		})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}

	//创建用户
	err := DB.Create(&newUser).Error // 在插入对象 newUser 必须带 & 一同传入,否则会报错 using unaddressable value

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
}

//查询手机号是否已存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	log.Println(user.ID) //输出 0
	db.Where("telephone = ?", telephone).First(&user)
	log.Println(user.ID)   //如果存在则输出 1,不存在打印 0
	log.Println(user.Name) //如果存在则输出对应的值 1,不存在打印 ""
	if user.ID != 0 {
		return true
	}
	return false
}
