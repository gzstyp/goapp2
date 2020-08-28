package route

import (
	"com.fwtai/app2/controller"
	"com.fwtai/app2/middleware"
	"github.com/gin-gonic/gin"
)

//定义路由列表
func CollectRoute(r *gin.Engine) *gin.Engine {
	//注册功能
	r.POST("/api/auth/register", controller.Register)
	//登录功能
	r.POST("/api/auth/login", controller.Login)
	//获取个人信息,使用 middleware.AuthMiddleware() 保护该接口,即需要认证的url接口
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
