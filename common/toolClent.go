package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//仅供参考
func createJsonExample(code int, msg string) map[string]interface{} {
	return gin.H{
		"code": 199,
		"msg":  "手机号必须是11位",
	}
}

//仅供参考
func responseJsonExample(context *gin.Context, json gin.H) {
	context.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": 199,
		"msg":  "手机号必须是11位",
	})
}

//包下的方法名,调用方式:common.CreateJson(199,"手机号必须是11位")
func CreateJson(code int, msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
}

//包下的方法名,调用方式:common.JsonData(199,"手机号必须是11位","token")
func JsonData(code int, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func ResponseJson(context *gin.Context, json map[string]interface{}) {
	context.JSON(http.StatusUnprocessableEntity, json)
}

func ResponseStatusJson(context *gin.Context, json map[string]interface{}, status int) {
	context.JSON(status, json)
}

//系统级的错误 50x
func CreateErrorJson(context *gin.Context, json map[string]interface{}) {
	context.JSON(http.StatusInternalServerError, json)
}
