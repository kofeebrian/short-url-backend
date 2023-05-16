package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kofeebrian/short-url-server/controllers"
	"github.com/kofeebrian/short-url-server/middlewares"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/token", controllers.GenerateToken)

	protected := r.Group("/", middlewares.AuthMiddleware)
	{
		protected.POST("/shorten", controllers.ShortenUrl)
	}

	return r
}
