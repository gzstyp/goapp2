package main

import (
	"com.fwtai/app2/common"
	"com.fwtai/app2/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := common.InitDB()
	defer db.Close() // 延时关闭它

	r := gin.Default()
	r = route.CollectRoute(r)

	panic(r.Run())
}
