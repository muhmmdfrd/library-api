package users

import (
	"net/http"

	"library-api/models/responses"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    userRepo UserRepo
}

func NewAuthHandler(userRepo UserRepo) *UserHandler {
    return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) Get(c *gin.Context) {
    users, err := h.userRepo.Get()
    if err != nil {
        c.JSON(http.StatusInternalServerError, responses.FailedResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, responses.SuccessResponse("success.", users))
}

func (h *UserHandler) Post(c *gin.Context) {
    success, err := h.userRepo.CreateTemp()
    if err != nil || !success {
        c.JSON(http.StatusBadRequest, responses.FailedResponse(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, responses.SuccessResponse("Data created.", nil))
}