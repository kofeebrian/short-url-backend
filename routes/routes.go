package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kofeebrian/short-url-server/controllers"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	group := r.Group("/api")
	{
		group.POST("/shorten", controllers.ShortenUrl)
	}

	return r
}
