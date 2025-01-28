package auth

import (
	middleware "library-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, handler *AuthHandler) {
    authRoutes := r.Group("/auth")
    {
        authRoutes.POST("/login", handler.Login)
        authRoutes.POST("/logout", middleware.AuthMiddleware(), handler.Logout)
    }
}