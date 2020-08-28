package route

import (
	"com.fwtai/app2/controller"
	"github.com/gin-gonic/gin"
)

//定义路由列表
func CollectRoute(r *gin.Engine) *gin.Engine {
	//注册功能
	r.POST("/api/auth/register", controller.Register)
	//登录功能
	r.POST("/api/auth/login", controller.Login)
	return r
}
