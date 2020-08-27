package route

import (
	"github.com/gin-gonic/gin"
	"goapp2/controller"
)

//在main.go调用
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
