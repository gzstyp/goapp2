package main

import (
	"com.fwtai/app2/common/database"
	"com.fwtai/app2/middleware"
	"com.fwtai/app2/route"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	initConfig() //初始化
	//初始化数据库
	db := database.InitDB()
	defer db.Close() // 延时关闭它
	//指定运行模式,默认是 gin.DebugMode,打包时需要把下面的这个注释取消掉!!!
	gin.SetMode(gin.DebugMode)
	e := gin.Default()

	// 不删除,前后端分离后用不到静态资源
	e.StaticFS("/static", http.Dir("./static"))             //加载整个静态文件目录,访问示例: http://192.168.3.108/static/win7.jpg
	e.StaticFile("/favicon.ico", "./resources/favicon.ico") //加载单个静态文件

	e.Use(cors()) // 开启跨域
	// 添加权限认证|权限拦截
	e.Use(middleware.AuthMiddleware())
	//注册controller路由接口
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
	wordDir, er := os.Getwd()
	if er != nil {
		log.Println("加载配置文件失败," + er.Error())
		return
	}
	viper.SetConfigName("application")       //文件名
	viper.SetConfigType("yml")               //文件类型
	viper.AddConfigPath(wordDir + "/config") //目录
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件出错," + err.Error())
	}
}

// 开启跨域处理
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		originDomains := []string{"http://192.168.3.108:8080", "http://localhost:8080"}
		inArraysFlag := false
		for _, value := range originDomains {
			if origin == value {
				inArraysFlag = true
				break
			}
		}
		if origin != "" && inArraysFlag {
			c.Header("Access-Control-Allow-Origin", origin)                                    // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options 探测请求")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
