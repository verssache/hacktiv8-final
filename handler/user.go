package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/hacktiv8-final4/auth"
	"github.com/verssache/hacktiv8-final4/helper"
	"github.com/verssache/hacktiv8-final4/user"
	"net/http"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := user.FormatRegisterUser(newUser)
	c.JSON(http.StatusCreated, formatter)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	loggedInUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := user.FormatLoginUser(token)
	c.JSON(http.StatusOK, formatter)
}

func (h *userHandler) TopUpBalance(c *gin.Context) {
	var input user.TopUpBalanceInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	total, err := h.userService.TopUpBalance(userID, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := user.FormatTopUpBalance(total)
	c.JSON(http.StatusOK, formatter)
}
