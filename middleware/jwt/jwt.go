package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"zrWorker/global"
	"zrWorker/pkg/e"
	"zrWorker/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("Authorization")

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if claims == nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "cookie失效，请点击右上角退出重新登陆",
					"data": data,
				})
				c.Abort()
				return
			}
			if err != nil {
				code = e.ERROR
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func Open() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		//slog.Printf(slog.INFO, "Authorization %s  %s", token, global.ServerSetting.BasicAuth)
		if token != global.ServerSetting.BasicAuth {
			code = e.ERROR
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
