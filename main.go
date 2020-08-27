package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goapp2/common"
	"goapp2/route"
)

func main() {

	db := common.InitDB()
	defer db.Close() // 延时关闭它

	r := gin.Default()
	r = route.CollectRoute(r)

	panic(r.Run())
}
