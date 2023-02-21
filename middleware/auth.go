package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/banking-api/util"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			data *util.JWTClaims
			err  error
		)
		cookie, err := c.Cookie("JWT")
		if err != nil {
			authHeader := c.Request.Header.Get("Authorization")
			if authHeader == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Authorization header empty.",
				})
				c.Abort()
				return
			}
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Authorization header format incorrect",
				})
				c.Abort()
				return
			}
			data, err = util.ParseToken(parts[1])
		} else {
			data, err = util.ParseToken(cookie)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
		c.Set("email", data.Email)
		c.Next()
	}

}
