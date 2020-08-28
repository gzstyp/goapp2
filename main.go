package main

import (
	"com.fwtai/app2/common/database"
	"com.fwtai/app2/middleware"
	"com.fwtai/app2/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	initConfig() //初始化
	//初始化数据库
	db := database.InitDB()
	defer db.Close() // 延时关闭它
	//指定运行模式,默认是 gin.DebugMode,打包时需要把下面的这个注释取消掉!!!
	gin.SetMode(gin.DebugMode)
	e := gin.Default()
	// 添加权限认证|权限拦截
	e.Use(middleware.AuthMiddleware())
	//注册接口
	e = route.CollectRoute(e)
	//_ = http.ListenAndServe(":80", e) //此方式没有Debug模式,适用于生产环境
	port := viper.GetString("server.port")
	if port != "" {
		_ = e.Run(":" + port)
	} else {
		_ = e.Run() // 也是ok的,Run()最终走的也是上面的 ListenAndServe
	}
}

//集中化管理配置
func initConfig() {
	wordDir, _ := os.Getwd()
	viper.SetConfigName("application")       //文件名
	viper.SetConfigType("yml")               //文件类型
	viper.AddConfigPath(wordDir + "/config") //目录
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件出错," + err.Error())
	}
}
