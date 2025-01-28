package auth

import (
	"library-api/models/requests"
	"library-api/modules/users"
	"library-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    userRepo users.UserRepo
}

func NewAuthHandler(userRepo users.UserRepo) *AuthHandler {
    return &AuthHandler{userRepo: userRepo}
}

func (h *AuthHandler) Login(c *gin.Context) {
    var request requests.AuthRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    user, err := h.userRepo.Auth(request)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := utils.GenerateToken(uint(user.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Simpan session di Redis
    if err := utils.StoreSession(c.Request.Context(), uint(user.ID), token); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store session"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
    userID := c.GetUint("user_id")
    if err := utils.DeleteSession(c.Request.Context(), userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}