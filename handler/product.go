package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/hacktiv8-final4/auth"
	"github.com/verssache/hacktiv8-final4/category"
	"github.com/verssache/hacktiv8-final4/helper"
	"github.com/verssache/hacktiv8-final4/product"
	"github.com/verssache/hacktiv8-final4/user"
	"net/http"
)

type productHandler struct {
	categoryService category.Service
	productService  product.Service
	authService     auth.Service
}

func NewProductHandler(categoryService category.Service, productService product.Service, authService auth.Service) *productHandler {
	return &productHandler{categoryService, productService, authService}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.CreateProductInput
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

	newProduct, err := h.productService.Store(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := product.FormatCreateProduct(newProduct)
	c.JSON(http.StatusCreated, formatter)
}

func (h *productHandler) FindAllProduct(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	if userID == 0 {
		errorMessage := gin.H{"errors": "Unauthorized"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	products, err := h.productService.FindAll()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := product.FormatGetAllProduct(products)
	c.JSON(http.StatusOK, formatter)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var input product.CreateProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	var productID product.FindProductInput
	err = c.ShouldBindUri(&productID)
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

	updatedProduct, err := h.productService.Update(productID.ID, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := product.FormatUpdateProduct(updatedProduct)
	c.JSON(http.StatusOK, formatter)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	var productID product.FindProductInput
	err := c.ShouldBindUri(&productID)
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

	err = h.productService.Delete(productID.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := product.FormatDeleteProduct()
	c.JSON(http.StatusOK, formatter)
}
