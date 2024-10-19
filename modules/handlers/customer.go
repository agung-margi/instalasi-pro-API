package handlers

import (
	helper "instalasi-pro/helpers"
	"instalasi-pro/middleware/auth"
	"instalasi-pro/modules/customer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerService customer.Service
	authService     auth.Service
}

func NewCustomerHandler(service customer.Service, authService auth.Service) *customerHandler {
	return &customerHandler{service, authService}
}

func (h *customerHandler) GetCustomers(c *gin.Context) {
	var input customer.CustomerInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get customers", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
}
