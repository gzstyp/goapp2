package controller

import (
	"com.fwtai/app2/common/database"
	"com.fwtai/app2/common/jwt"
	"com.fwtai/app2/common/toolClent"
	"com.fwtai/app2/dto"
	"com.fwtai/app2/model"
	"com.fwtai/app2/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

//注册
func Register(context *gin.Context) {
	DB := database.GetDB() //前面的 DB 写成db小写也是可以的,下面要跟着改
	//1.获取参数
	name := context.PostForm("name")
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")

	//2.数据验证
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
	//3.判断手机号是否已存在
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
	hashPassword, er := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if er != nil {
		toolClent.ResponseJson(context, toolClent.JsonErrorMsg("加密出错"))
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}

	//4.创建用户
	err := DB.Create(&newUser).Error // 在插入对象 newUser 必须带 & 一同传入,否则会报错 using unaddressable value

	if err != nil {
		context.JSON(200, gin.H{
			"msg": "注册失败," + err.Error(),
		})
		return
	}
	//5.返回结果
	toolClent.ResponseJson(context, toolClent.JsonSucceedMsg("注册成功,手机号:"+telephone))
}

//登录
func Login(context *gin.Context) {
	//1.获取参数
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")
	//2.数据验证
	if len(telephone) != 11 {
		toolClent.Response199Msg(context, "请输入正确的手机号")
		return
	}
	if len(password) < 6 {
		toolClent.Response199Msg(context, "密码长度需要大于6位")
		return
	}
	//3.判断手机号和密码是否存在
	db := database.GetDB()
	var user model.User
	//err := db.Where("telephone = ? and password = ?", telephone,password).First(&user).Error //多参数,好使的!!!
	err := db.Where("telephone = ?", telephone).First(&user).Error
	if err != nil {
		msg := err.Error()
		if msg == "record not found" {
			toolClent.Response199Msg(context, "用户名或密码错误!")
			return
		}
		toolClent.Response204Msg(context, "系统出现错误,"+err.Error())
		return
	}
	if user.ID == 0 {
		toolClent.Response199Msg(context, "用户名或密码错误")
		return
	}
	if e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); e != nil {
		toolClent.Response199Msg(context, "用户名或密码错误")
		return
	}
	//4.发放token
	token, eToken := jwt.CreateToken(user)
	if eToken != nil {
		toolClent.Response204Msg(context, "哦豁,生成token错误!")
		log.Printf("生成token错误:%v", eToken)
		return
	}
	data := gin.H{
		"token":    token,
		"menuData": "[{}]",
	}
	//5.返回结果
	toolClent.ResponseData(context, "登录成功", data)
}

//获取个人信息(需要权限认证才能访问的)
func Info(context *gin.Context) {
	//获取个人信息肯定是通过认证了，即肯定 context.Set("user", user)，所以直接从 context 获取
	user, _ := context.Get("user")              // _ 用不到该值时就用占位符
	userDto := dto.ToUserDto(user.(model.User)) // 强制类型转换???
	toolClent.ResponseData(context, "操作成功", userDto)
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
