package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/pkg/e"
	"goblog/utils"
	"net/http"
	"strings"
)

// JWT jwt中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = e.ErrorTokenExist //1004
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": e.GetMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": e.GetMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": e.GetMsg(code),
			})
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == utils.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  e.ERROR,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  e.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims)
		c.Next()
	}
}
