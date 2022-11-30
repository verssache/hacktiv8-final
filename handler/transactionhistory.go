package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/verssache/hacktiv8-final4/auth"
	"github.com/verssache/hacktiv8-final4/category"
	"github.com/verssache/hacktiv8-final4/helper"
	"github.com/verssache/hacktiv8-final4/product"
	"github.com/verssache/hacktiv8-final4/transactionhistory"
	"github.com/verssache/hacktiv8-final4/user"
	"net/http"
)

type transactionHistoryHandler struct {
	transactionHistoryService transactionhistory.Service
	userService               user.Service
	productService            product.Service
	categoryService           category.Service
	authService               auth.Service
}

func NewTransactionHistoryHandler(transactionHistoryService transactionhistory.Service, userService user.Service, productService product.Service, categoryService category.Service, authService auth.Service) *transactionHistoryHandler {
	return &transactionHistoryHandler{transactionHistoryService, userService, productService, categoryService, authService}
}

func (h *transactionHistoryHandler) CreateTransaction(c *gin.Context) {
	var input transactionhistory.CreateTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	userBalance := currentUser.Balance

	if userID == 0 {
		errorMessage := gin.H{"errors": "Unauthorized"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	product, err := h.productService.FindByID(input.ProductID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	if product.ID == 0 {
		errorMessage := gin.H{"errors": "Product not found"}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	if product.Stock < input.Quantity {
		errorMessage := gin.H{"errors": "Stock is not enough"}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	totalPrice := product.Price * input.Quantity

	if userBalance < totalPrice {
		errorMessage := gin.H{"errors": "Balance is not enough"}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	_, err = h.productService.UpdateStock(product.ID, input.Quantity)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	_, err = h.userService.UpdateBalance(userID, totalPrice)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	_, err = h.categoryService.UpdateSoldProduct(product.CategoryID, input.Quantity)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	newTransactionHistory, err := h.transactionHistoryService.Store(input, userID, totalPrice)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := transactionhistory.FormatCreateTransactionHistory(newTransactionHistory, product.Title)
	c.JSON(http.StatusCreated, formatter)
}

func (h *transactionHistoryHandler) GetMyTransactionHistory(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	if userID == 0 {
		errorMessage := gin.H{"errors": "Unauthorized"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	transactionHistories, err := h.transactionHistoryService.FindAllByID(userID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := transactionhistory.FormatMyTransaction(transactionHistories)
	c.JSON(http.StatusOK, formatter)
}

func (h *transactionHistoryHandler) GetAllTransactionHistory(c *gin.Context) {
	currentUserRole := c.MustGet("currentUserRole").(string)

	if currentUserRole != "admin" {
		errorMessage := gin.H{"errors": "Unauthorized"}
		c.JSON(http.StatusUnauthorized, errorMessage)
		return
	}

	transactionHistories, err := h.transactionHistoryService.FindAll()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := transactionhistory.FormatAllTransactionHistory(transactionHistories)
	c.JSON(http.StatusOK, formatter)
}
