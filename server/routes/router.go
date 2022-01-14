package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/PedroPaiao/rest-api-go/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.IndexBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DestroyBook)
		}
	}

	return router
}
