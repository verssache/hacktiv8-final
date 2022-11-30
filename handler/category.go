package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/hacktiv8-final4/auth"
	"github.com/verssache/hacktiv8-final4/category"
	"github.com/verssache/hacktiv8-final4/helper"
	"github.com/verssache/hacktiv8-final4/user"
	"net/http"
)

type categoryHandler struct {
	categoryService category.Service
	authService     auth.Service
}

func NewCategoryHandler(categoryService category.Service, authService auth.Service) *categoryHandler {
	return &categoryHandler{categoryService, authService}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CreateCategoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	currentRole := c.MustGet("currentUserRole").(string)
	if currentRole != "admin" {
		errorMessage := gin.H{"errors": "You are not authorized to create category"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	newCategory, err := h.categoryService.Store(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := category.FormatCreateCategory(newCategory)
	c.JSON(http.StatusCreated, formatter)
}

func (h *categoryHandler) FindAllCategory(c *gin.Context) {
	currentRole := c.MustGet("currentUserRole").(string)
	if currentRole != "admin" {
		errorMessage := gin.H{"errors": "You are not authorized to create category"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	categories, err := h.categoryService.FindAll()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := category.FormatGetAllCategory(categories)
	c.JSON(http.StatusOK, formatter)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var input category.CreateCategoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	var inputID category.FindCategoryInput
	err = c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	currentRole := c.MustGet("currentUserRole").(string)
	if currentRole != "admin" {
		errorMessage := gin.H{"errors": "You are not authorized to create category"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	updatedCategory, err := h.categoryService.Update(inputID.ID, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := category.FormatCreateCategory(updatedCategory)
	c.JSON(http.StatusOK, formatter)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	var inputID category.FindCategoryInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentRole := c.MustGet("currentUserRole").(string)
	if currentRole != "admin" {
		errorMessage := gin.H{"errors": "You are not authorized to create category"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	err = h.categoryService.Delete(inputID.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := category.FormatDeleteCategory()
	c.JSON(http.StatusOK, formatter)
}
