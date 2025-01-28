package auth

import (
	"library-api/models/requests"
	"library-api/models/responses"
	"library-api/modules/users"
	"library-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    userService users.UserService
}

func NewAuthHandler(userService users.UserService) *AuthHandler {
    return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Login(c *gin.Context) {
    var request requests.AuthRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, responses.FailedResponse("Invalid input."))
        return
    }

    user, err := h.userService.Auth(request)
    if err != nil {
        c.JSON(http.StatusUnauthorized, responses.FailedResponse("Invalid credentials."))
        return
    }

    token, err := utils.GenerateToken(uint(user.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, responses.FailedResponse("Failed to generate token."))
        return
    }

    // Save session at Redis
    if err := utils.StoreSession(c.Request.Context(), uint(user.ID), token); err != nil {
        c.JSON(http.StatusInternalServerError, responses.FailedResponse("Failed to store session."))
        return
    }

    c.JSON(http.StatusOK, responses.SuccessResponse("success.", token))
}

func (h *AuthHandler) Logout(c *gin.Context) {
    userID := c.GetUint("user_id")
    if err := utils.DeleteSession(c.Request.Context(), userID); err != nil {
        c.JSON(http.StatusInternalServerError, responses.FailedResponse("Failed to delete session."))
        return
    }

    c.JSON(http.StatusOK, responses.SuccessResponse("logged out", nil))
}