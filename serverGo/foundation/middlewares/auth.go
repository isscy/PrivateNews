package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"serverGo/foundation/jwt"
	"serverGo/foundation/resp"
	"serverGo/foundation/utils"
)

/**
基于JWT的认证中间件
*/

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			resp.Fails(c, resp.InvalidAuthFormat, "缺失Token")
			c.Abort()
			return
		}
		// 分割获取Token
		/*parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2) {
			resp.Fails(c, resp.InvalidAuthFormat, "Token格式错误")
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(parts[1])*/
		claims, err := jwt.ParseToken(authHeader)
		if err != nil {
			fmt.Println(err)
			resp.Fail(c, resp.InvalidToken)
			c.Abort()
			return
		}
		c.Set(utils.ContextUserIdKey, claims.UserID) // 将当前请求的userID信息保存到请求的上下文c上
		c.Set(utils.ContextUserNameKey, claims.Username)
		c.Next()

	}

}
