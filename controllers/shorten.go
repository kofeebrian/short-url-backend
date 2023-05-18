package controllers

import (
	"log"
	"net/http"
	gourl "net/url"
	"os"

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
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Bad request", "error": err.Error()})
		return
	}

	err := models.CreateShortenedUrl(input.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error", "error": err.Error()})
		return
	}

	result, err := models.GetShortenedUrl(input.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Success", "result_url": result.ShortenedUrl})
}

func RedirectUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	shortUrl, err := gourl.JoinPath(os.Getenv("SERVER_NAME"), shortUrl)
	log.Printf("shortUrl: %s", shortUrl)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Bad request", "error": err.Error()})
		return
	}

	result, err := models.GetOriginalUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Url not found", "error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, result.LongUrl)
}
