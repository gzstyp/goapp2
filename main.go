package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main(){
	r := gin.Default()
	r.POST("/api/auth/register", func(context *gin.Context) {
		//获取参数
		name := context.PostForm("name")
		telephone := context.PostForm("telephone")
		password := context.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			context.JSON(http.StatusUnprocessableEntity,gin.H{
				"code" : 199,
				"msg" : "手机号必须是11位",
			})
			return
		}
		if len(password) < 6 {
			context.JSON(422,map[string]interface{}{
				"code" : 199,
				"msg" : "密码必须大于6位",
			})
			return
		}
		//判断手机号
		if telephone == "13765121695" {
			context.JSON(422,gin.H{
				"code" : 199,
				"msg" : "手机号已被占用",
			})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name,telephone,password)
		//创建用户
		//返回结果
		context.JSON(200,gin.H{
			"msg" : "注册成功,手机号:"+telephone,
		})
	})
	panic(r.Run())
}

//随机生成10位的字符串
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnm123456987MNBVCXZASDFGHJKLPOIUYTREWQ")
	result := make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}