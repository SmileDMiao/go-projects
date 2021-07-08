package middleware

import (
	"go-learn/global"
	"go-learn/model"
	"go-learn/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// JWTAuth handler
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{"message": "Token不存在"})
			c.Abort()
			return
		}
		j := service.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		var user model.User
		result := global.DB.Where("`id` = ?", claims.ID).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{"message": "用户不存在"})
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
