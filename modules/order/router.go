package order

import (
	"fmt"
	"instalasi-pro/database"
	helper "instalasi-pro/helpers"
	"instalasi-pro/middleware/auth"
	"instalasi-pro/modules/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/orders")
	api.Use(auth.AuthMiddleware())
	{
		api.GET("/", FindAll)
		api.GET("/:id", FindById)
		api.GET("/user/:id", FindByUserID)
		api.POST("/", Save)
		// api.PUT("/:id", Update)
	}
}

func FindAll(c *gin.Context) {
	orders, err := NewService(NewRepository(database.DB)).FindAll()
	if err != nil {
		response := helper.APIResponse("Failed to get orders", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of orders", http.StatusOK, "success", orders)
	c.JSON(http.StatusOK, response)
}

func FindById(c *gin.Context) {
	productId := c.Param("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		response := helper.APIResponse("Failed to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := NewService(NewRepository(database.DB)).FindById(id)
	if err != nil {
		response := helper.APIResponse("Failed to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of product", http.StatusOK, "success", order)
	c.JSON(http.StatusOK, response)
}

func FindByUserID(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := NewService(NewRepository(database.DB)).FindByUserID(id)
	if err != nil {
		response := helper.APIResponse("Failed to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of product", http.StatusOK, "success", order)
	c.JSON(http.StatusOK, response)
}

func Save(c *gin.Context) {
	if isCustomer, err := CheckIfCustomer(c); err != nil || !isCustomer {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	var input OrderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	order, err := NewService(NewRepository(database.DB)).CreateOrder(input)

	if err != nil {
		response := helper.APIResponse("Failed to create order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create order", http.StatusOK, "success", order)
	c.JSON(http.StatusOK, response)
}

func CheckIfCustomer(c *gin.Context) (bool, error) {
	currentUser, exists := c.MustGet("currentUser").(user.User)

	fmt.Println(currentUser)
	if !exists {
		return false, fmt.Errorf("User not authenticated")
	}

	if currentUser.Role != "customer" {
		return false, fmt.Errorf("Only customers can create orders")
	}

	return true, nil
}
