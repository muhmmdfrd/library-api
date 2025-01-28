package middlewares

import (
	"library-api/models/responses"
	"library-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, responses.FailedResponse("Authorization header required."))
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, responses.FailedResponse("Invalid token."))
            c.Abort()
            return
        }

        sessionToken, err := utils.GetSession(c.Request.Context(), claims.UserID)
        if err != nil || sessionToken != tokenString {
            c.JSON(http.StatusUnauthorized, responses.FailedResponse("Invalid session."))
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Next()
    }
}