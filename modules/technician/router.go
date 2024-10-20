package technician

import (
	"instalasi-pro/database"
	helper "instalasi-pro/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/technicians")
	{
		api.GET("/", FindTechnicians)
		api.GET("/:id", FindByTechnicianID)
		api.POST("/", RegisterTechnician)

	}
}

func FindTechnicians(c *gin.Context) {
	technicians, err := NewService(NewRepository(database.DB)).FindAll()
	if err != nil {
		response := helper.APIResponse("Failed to get technicians", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of technicians", http.StatusOK, "success", technicians)
	c.JSON(http.StatusOK, response)

}

func FindByTechnicianID(c *gin.Context) {
	technicianID := c.Param("id")

	id, err := strconv.Atoi(technicianID)
	if err != nil {
		response := helper.APIResponse("Failed to get technician", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	technician, err := NewService(NewRepository(database.DB)).FindById(id)
	if err != nil {
		response := helper.APIResponse("Failed to get technician", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if technician.Role != "technician" {
		response := helper.APIResponse("Technician not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Detail of technician", http.StatusOK, "success", technician)
	c.JSON(http.StatusOK, response)
}

func RegisterTechnician(c *gin.Context) {
	var input TechnicianInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userRepo := NewRepository(database.DB)
	userService := NewService(userRepo)

	newUser, err := userService.RegisterTechnician(input)
	if err != nil {
		response := helper.APIResponse("Register failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success", http.StatusOK, "success", newUser)
	c.JSON(http.StatusOK, response)
}
