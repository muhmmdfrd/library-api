package users

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, handler *UserHandler) {
  userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", handler.Get)
        userRoutes.POST("/", handler.Post)
    }
    
}