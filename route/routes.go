package route

import (
	"com.fwtai/app2/controller"
	"github.com/gin-gonic/gin"
)

//定义路由列表
func CollectRoute(e *gin.Engine) *gin.Engine {
	e.GET("/", controller.Index)
	//注册功能
	e.POST("/api/auth/register", controller.Register)
	//登录功能
	e.POST("/api/auth/login", controller.Login)
	//获取个人信息
	e.GET("/api/auth/info", controller.Info)
	// 如果有少量的url接口需要被保护的,可以使用 middleware.AuthMiddleware() 保护该接口,即需要认证的url接口[需要权限认证才能访问的]
	//e.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	e.GET("/redire", controller.Redire)
	return e
}
