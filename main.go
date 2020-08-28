package main

import (
	"com.fwtai/app2/common/database"
	"com.fwtai/app2/middleware"
	"com.fwtai/app2/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//初始化数据库
	db := database.InitDB()
	defer db.Close() // 延时关闭它
	//指定运行模式,默认是 gin.DebugMode,打包时需要把下面的这个注释取消掉!!!
	//gin.SetMode(gin.ReleaseMode)
	e := gin.Default()
	// 添加权限认证|权限拦截
	e.Use(middleware.AuthMiddleware())
	//注册接口
	e = route.CollectRoute(e)
	//_ = http.ListenAndServe(":80", e) //此方式没有Debug模式,适用于生产环境,
	_ = e.Run(":80") // 也是ok的,最终走的也是上面的 ListenAndServe,有Debug模式|生产环境模式
}
