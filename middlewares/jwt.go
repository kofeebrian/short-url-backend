package middlewares

import (
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	s := c.GetHeader("Authorization")
	if s == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized", "error": "No Authorization header"})
		return
	}

	tokenString := strings.TrimPrefix(s, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized", "error": err.Error()})
		return
	}

	if _, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized", "error": err.Error()})
		return
	}
}
