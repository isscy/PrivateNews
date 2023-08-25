package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
 jwt包自带的jwt.StandardClaims只包含了官方字段
我们这里需要额外记录一个UserID字段，所以要自定义结构体
自定义声明结构体并内嵌jwt.StandardClaims
*/

type MyClaims struct {
	UserID   uint64 `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义Secret 用于加密的字符串
var mySecret = []byte("Private-News-Server-Go-Secret")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// 定义JWT的过期时间
const TokenExpireDuration = time.Hour * 24
const AccessTokenExpireDuration = time.Hour * 24      // access_token 过期时间
const RefreshTokenExpireDuration = time.Hour * 24 * 7 // refresh_token 过期时间

// 生成token
func GenToken(userId uint64, username string) (accessToken, refreshToken string, err error) {
	c := MyClaims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "SysDefault",                                     // 签发人
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // 过期时间
		Issuer:    "SysDefault",                                      // 签发人
	}).SignedString(mySecret)
	return
}

// ParseToken 解析 token/*
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("无效的token")
	}
	return
}

func RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	if _, err = jwt.Parse(refreshToken, keyFunc); err != nil {
		return
	}
	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims MyClaims
	_, err = jwt.ParseWithClaims(accessToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired { // 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
