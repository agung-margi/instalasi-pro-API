package product

import (
	"instalasi-pro/database"
	helper "instalasi-pro/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/products")
	{
		api.GET("/", FindAll)
		api.GET("/:id", FindById)
		api.POST("/", Save)
		api.PUT("/:id", UpdateStatus)
	}
}

func FindAll(c *gin.Context) {
	products, err := NewService(NewRepository(database.DB)).FindAll()
	if err != nil {
		response := helper.APIResponse("Failed to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", products)
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
	product, err := NewService(NewRepository(database.DB)).FindById(id)
	if err != nil {
		response := helper.APIResponse("Failed to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func Save(c *gin.Context) {
	var input CreateProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	product, err := NewService(NewRepository(database.DB)).Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to create product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func UpdateStatus(c *gin.Context) {
	productId := c.Param("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		response := helper.APIResponse("Failed to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productService := NewService(NewRepository(database.DB))
	updatedProduct, err := productService.UpdateStatus(id)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update product", http.StatusOK, "success", updatedProduct)
	c.JSON(http.StatusOK, response)
}
