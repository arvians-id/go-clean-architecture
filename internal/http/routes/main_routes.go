package routes

import (
	"github.com/arvians-id/go-clean-architecture/internal/http/controller"
	"github.com/arvians-id/go-clean-architecture/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewInitializedRoutes(controller controller.UserController) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.SetupCorsMiddleware())

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", controller.FindAll)
			users.GET("/:id", controller.FindByID)
			users.POST("/", controller.Create)
			users.PATCH("/:id", controller.Update)
			users.DELETE("/:id", controller.Delete)
		}
	}

	return router
}
