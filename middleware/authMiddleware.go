package middleware

import (
	"com.fwtai/app2/common/database"
	"com.fwtai/app2/common/jwt"
	"com.fwtai/app2/common/toolClent"
	"com.fwtai/app2/model"
	"github.com/gin-gonic/gin"
	"strings"
)

//中间件的使用，在API中可能使用限流|身份验证等;本项目是用于认证token,即需要权限认证都要添加本方法,调用方式: e.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		url := context.Request.RequestURI
		if auth(url) {
			context.Next()
			return
		}
		//1.获取authorization header
		tokenString := context.GetHeader("Authorization")
		//2.验证格式 validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			toolClent.Response401(context)
			context.Abort() //将本次请求终止|阻止|结束本次请求
			return
		}
		//截取token
		tokenString = tokenString[7:] //从第0到7个开始截取
		token, claims, err := jwt.ParseToken(tokenString)
		//如果解析失败或token无效,提示重新登录
		if err != nil || !token.Valid {
			toolClent.Response205Msg(context, "无效token或token已失效,请重新登录")
			context.Abort() //将本次请求终止|阻止|结束本次请求
			return
		}
		//3.验证通过后,从 claims 抽取userId,claims是个model实体
		userId := claims.UserId
		DB := database.GetDB()
		var user model.User
		DB.First(&user, userId)
		//验证用户是否存在
		if user.ID == 0 {
			toolClent.Response199Msg(context, "用户不存在")
			return
		}
		//若用户userId存在,则将user信息写入上下文
		context.Set("user", user)
		context.Next() //继续执行后面的逻辑
	}
}

//无需权限认证或不受保护的接口
func auth(url string) bool {
	var arrs = [...]string{"/", "/redire", "/single", "/multi", "/api/auth/register", "/api/auth/login", "/api/auth/logout"}
	for i := 0; i < len(arrs); i++ {
		if arrs[i] == url {
			return true
		}
	}
	return false
}
