package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"practice/pkg/config"
	"practice/src/user-service/token"
	"strings"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSONP(http.StatusUnauthorized, gin.H{
				"Message": "Request doest not contain token",
			})
			log.Println("MiddleWare: Error to get token in")
			c.Abort()
			return
		}
		tokenResult := strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := token.ExtractToken(tokenResult)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Invalid Token",
			})
			log.Println("Error: Invalid token")
			c.Abort()
			return
		}

		c.Set(config.UserId, claims.Id)
		c.Next()
	}
}
