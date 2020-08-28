package controller

import (
	"com.fwtai/app2/common"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	common.ResponseJson(context, common.CreateJson(200, "感谢使用go后端服务接口"))
}
