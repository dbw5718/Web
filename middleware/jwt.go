package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"todo_list/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		fmt.Println(token)
		//claims,_:=utils.ParseToken(token)
		//fmt.Println(claims)
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401
			}
		}
		//fmt.Println(code)
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
