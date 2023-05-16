package controllers

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    c.ClientIP(),
	})

	ss, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": ss})
}
