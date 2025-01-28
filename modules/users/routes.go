package users

import (
	"library-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, handler *UserHandler) {
  userRoutes := r.Group("/api/v1/users", middlewares.AuthMiddleware())
    {
        userRoutes.GET("/", handler.Get)
        userRoutes.POST("/", handler.Post)
    }
    
}