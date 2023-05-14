package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kofeebrian/short-url-server/models"
)

type ShortenedUrlInput struct {
	Url string `json:"url" binding:"required,url"`
}

func ShortenUrl(c *gin.Context) {

	// Get Request input
	var input ShortenedUrlInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.CreateShortenedUrl(input.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Success"})
}
