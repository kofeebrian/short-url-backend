package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kofeebrian/short-url-server/controllers"
	"github.com/kofeebrian/short-url-server/middlewares"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	r.GET("/token", controllers.GenerateToken)
	r.GET("/:shortUrl", controllers.RedirectUrl)

	protected := r.Group("/", middlewares.AuthMiddleware)
	{
		protected.POST("/shorten", controllers.ShortenUrl)
	}

	return r
}
