package route

import (
	"com.fwtai/app2/controller"
	"github.com/gin-gonic/gin"
)

//在main.go调用
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
