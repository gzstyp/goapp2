package toolClent

import (
	"com.fwtai/app2/common/configFile"
	"github.com/gin-gonic/gin"
	"net/http"
)

//生成json格式数据,包下的方法名,调用方式:toolClent.JsonData(configFile.Code199,"手机号必须是11位")
func JsonData(code int, msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
}

//生成操作失败的json数据格式
func JsonFail() map[string]interface{} {
	return map[string]interface{}{
		"code": configFile.Code199,
		"msg":  configFile.Msg199,
	}
}

//生成操作失败的json数据格式,可指定提示文字
func JsonFailMsg(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": configFile.Code199,
		"msg":  msg,
	}
}

//生成操作成功的json数据格式
func JsonSucceed() map[string]interface{} {
	return map[string]interface{}{
		"code": configFile.Code200,
		"msg":  configFile.Msg200,
	}
}

//生成操作成功的json数据格式,可指定提示文字
func JsonSucceedMsg(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": configFile.Code200,
		"msg":  msg,
	}
}

//生成系统错误的json数据格式
func JsonError() map[string]interface{} {
	return map[string]interface{}{
		"code": configFile.Code204,
		"msg":  configFile.Msg204,
	}
}

//生成系统错误的json数据格式,可指定提示文字
func JsonErrorMsg(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": configFile.Code204,
		"msg":  msg,
	}
}

// 响应客户端,增删改操作调用
func ResponseExecute(context *gin.Context, bl bool, msg string) {
	json := map[string]interface{}{
		"code": configFile.Code200,
		"msg":  msg,
	}
	if !bl {
		json = map[string]interface{}{
			"code": configFile.Code199,
			"msg":  configFile.Msg199,
		}
	}
	context.JSON(http.StatusOK, json)
}

//响应客户端,数据类型:json数据格式
func ResponseJson(context *gin.Context, json map[string]interface{}) {
	context.JSON(http.StatusOK, json)
}

//响应客户端,code=200,数据类型:不限类型数据格式
func ResponseObj(context *gin.Context, data interface{}) {
	payload := map[string]interface{}{
		"code": configFile.Code200,
		"msg":  configFile.Msg200,
		"data": data,
	}
	context.JSON(http.StatusOK, payload)
}

//响应客户端,数据类型:不限类型数据格式
func ResponseData(context *gin.Context, msg string, data interface{}) {
	payload := map[string]interface{}{
		"code": configFile.Code200,
		"msg":  msg,
		"data": data,
	}
	if data == nil || data == "" {
		payload = map[string]interface{}{
			"code": configFile.Code201,
			"msg":  configFile.Msg201,
		}
	}
	context.JSON(http.StatusOK, payload)
}

//响应客户端,操作失败
func Response199(context *gin.Context) {
	payload := map[string]interface{}{
		"code": configFile.Code199,
		"msg":  configFile.Msg199,
	}
	context.JSON(http.StatusOK, payload)
}

//响应客户端,操作失败,可指定提示文字
func Response199Msg(context *gin.Context, msg string) {
	payload := map[string]interface{}{
		"code": configFile.Code199,
		"msg":  msg,
	}
	context.JSON(http.StatusOK, payload)
}

//响应客户端,系统出现错误
func Response204(context *gin.Context) {
	payload := map[string]interface{}{
		"code": configFile.Code204,
		"msg":  configFile.Msg204,
	}
	context.JSON(http.StatusOK, payload)
}

//响应客户端,系统出现错误,可指定提示文字
func Response204Msg(context *gin.Context, msg string) {
	payload := map[string]interface{}{
		"code": configFile.Code204,
		"msg":  msg,
	}
	context.JSON(http.StatusOK, payload)
}

//响应客户端,操作成功
func Response200(context *gin.Context) {
	json := map[string]interface{}{
		"code": configFile.Code200,
		"msg":  configFile.Msg200,
	}
	context.JSON(http.StatusOK, json)
}

//响应客户端,操作成功,可指定提示文字
func Response200Msg(context *gin.Context, msg string) {
	json := map[string]interface{}{
		"code": configFile.Code200,
		"msg":  msg,
	}
	context.JSON(http.StatusOK, json)
}

//没有操作权限|权限不足
func Response401(context *gin.Context) {
	json := map[string]interface{}{
		"code": configFile.Code401,
		"msg":  configFile.Msg401,
	}
	context.JSON(http.StatusOK, json)
}

//没有操作权限|权限不足,可指定提示文字
func Response401Msg(context *gin.Context, msg string) {
	json := map[string]interface{}{
		"code": configFile.Code401,
		"msg":  msg,
	}
	context.JSON(http.StatusOK, json)
}

//没有操作权限|权限不足
func Response205(context *gin.Context) {
	json := map[string]interface{}{
		"code": configFile.Code205,
		"msg":  configFile.Msg205,
	}
	context.JSON(http.StatusOK, json)
}

//没有操作权限|权限不足,可指定提示文字
func Response205Msg(context *gin.Context, msg string) {
	json := map[string]interface{}{
		"code": configFile.Code205,
		"msg":  msg,
	}
	context.JSON(http.StatusOK, json)
}
