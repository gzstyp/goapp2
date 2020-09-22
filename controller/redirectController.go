package controller

import (
	"com.fwtai/app2/common/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

//重定向
func Redire(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, "http://www.yinlz.com")
	database.InitSqlServer()
}
