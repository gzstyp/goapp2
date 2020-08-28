package controller

import (
	"com.fwtai/app2/common/toolClent"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	toolClent.Response200Msg(context, "感谢使用go后端服务接口")
}
