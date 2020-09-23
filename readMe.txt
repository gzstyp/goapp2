
https://goproxy.cn

go get -u github.com/gin-gonic/gin

#Mysql数据库
go get -u github.com/jinzhu/gorm

#MySQL数据库驱动
go get -u github.com/go-sql-driver/mysql

#加密
go get -u golang.org/x/crypto/bcryp

#jwt
go get github.com/dgrijalva/jwt-go

#config读取配置文件
go get github.com/spf13/viper

查看|依赖|安装依赖(可选)
go mod tidy

运行
go run main.go ./

生成执行文件
go build

运行执行文件
goapp2.exe

windows下打包部署到Linux

set GOARCH=amd64
set GOOS=linux
go build main.go

推荐
go build -ldflags "-w -s" main.go

给文件添加执行的权限
chmod +x main

创建目录并上传文件配置文件 application.yml
mkdir -p config

运行
./main
