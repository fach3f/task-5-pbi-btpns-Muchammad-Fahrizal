package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"tugasakhir/helpers"
)

func Auth(c *gin.Context) {
	tokenStr, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(helpers.JWT_SECRET), nil
	})

	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userid", claims["sub"])
		c.Next()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JWT token",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
