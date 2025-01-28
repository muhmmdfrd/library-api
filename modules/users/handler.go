package users

import (
	"library-api/modules/users/models"
	"library-api/utils"
	"net/http"

	"library-api/models/responses"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) Get(c *gin.Context) {
    filter, err := utils.ParseFilter(c)
    if err != nil {
        c.JSON(http.StatusBadRequest, responses.FailedResponse(err.Error()))
        return
    }
    
    users, err := h.userService.GetPagedUser(filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, responses.FailedResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, responses.SuccessResponse("success.", users))
}

func (h *UserHandler) Post(c *gin.Context) {
    var add models.UserAdd

	err := c.ShouldBindJSON(&add)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.FailedResponse(err.Error()))
        return
	}

    view, err := h.userService.CreateUser(add)
    if err != nil {
        c.JSON(http.StatusBadRequest, responses.FailedResponse(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, responses.SuccessResponse("Data created.", view))
}