package user

import (
	"instalasi-pro/database"
	helper "instalasi-pro/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/users")
	{
		api.POST("/register", Register)
		api.POST("/login", Login)
		api.PUT("/:id", UpdateUser)
		api.GET("/:id", FindByID)
		api.GET("/", FindAll)
	}
}

func Register(c *gin.Context) {
	var input UserInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userRepo := NewRepository(database.DB)
	userService := NewUserService(userRepo)

	newUser, err := userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatRegister(newUser)

	response := helper.APIResponse("success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func Login(c *gin.Context) {
	var input UserInput
	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userService := NewUserService(NewRepository(database.DB))
	token, err := userService.Login(input)

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Login success", http.StatusOK, "success", gin.H{"token": token})
	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		response := helper.APIResponse("Invalid user ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input UpdateUserInput
	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userService := NewUserService(NewRepository(database.DB))
	newUser, err := userService.UpdateDataUser(id, input)
	if err != nil {
		response := helper.APIResponse("Update user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := FormatUser(newUser)

	response := helper.APIResponse("success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func FindByID(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		response := helper.APIResponse("Invalid user ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userService := NewUserService(NewRepository(database.DB))
	user, err := userService.GetUser(id)
	if err != nil {
		response := helper.APIResponse("Get user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := FormatUsers(user)

	response := helper.APIResponse("success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func FindAll(c *gin.Context) {
	userService := NewUserService(NewRepository(database.DB))
	users, err := userService.GetUser(0)
	if err != nil {
		response := helper.APIResponse("Get users failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := FormatUsers(users)
	response := helper.APIResponse("success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
