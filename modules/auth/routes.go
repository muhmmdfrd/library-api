package auth

import (
	middleware "library-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, handler *AuthHandler) {
    authRoutes := r.Group("/api/v1/auth")
    {
        authRoutes.POST("/", handler.Login)
        authRoutes.POST("/logout", middleware.AuthMiddleware(), handler.Logout)
    }
}