package jwt

import (
	"com.fwtai/app2/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//加密的密钥|盐值
var jwtKey = []byte("a_secret_crect")

//model|结构体|java的实体
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//生成token,方法的定义!!!
func CreateToken(user model.User) (string, error) {
	//设置token的过期时间,7天 ==> time.Hour * 24 * 7
	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "引路者|www.yinlz.com",   //发放人
			Subject:   "用户token",             //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //密钥加密
	//如果生成错误,将错误返回
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
