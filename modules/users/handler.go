package users

import (
	"bytes"
	"encoding/json"
	"net/http"

	"library-api/models/responses"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    userRepo UserRepo
}

func NewUserHandler(userRepo UserRepo) *UserHandler {
    return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) Get(c *gin.Context) {
    users, err := h.userRepo.Get()
    if err != nil {
        c.JSON(http.StatusInternalServerError, responses.FailedResponse(err.Error()))
        return
    }

    var buf bytes.Buffer
    encoder := json.NewEncoder(&buf)
    encoder.Encode(responses.SuccessResponse("success.", users))

    c.Data(http.StatusOK, "application/json", buf.Bytes())
}

func (h *UserHandler) Post(c *gin.Context) {
    success, err := h.userRepo.CreateTemp()
    if err != nil || !success {
        c.JSON(http.StatusBadRequest, responses.FailedResponse(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, responses.SuccessResponse("Data created.", nil))
}