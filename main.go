package main

import (
	"instalasi-pro/configs"
	"instalasi-pro/database"
	"instalasi-pro/modules/invoice"
	"instalasi-pro/modules/order"
	"instalasi-pro/modules/product"
	"instalasi-pro/modules/technician"
	"instalasi-pro/modules/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()
	database.Connection()

	database.DB.AutoMigrate(&user.User{})
	database.DB.AutoMigrate(&product.Product{})
	database.DB.AutoMigrate(&order.Order{})
	database.DB.AutoMigrate(&invoice.Invoice{})

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Instalasi-Pro",
		})
	})
	user.Initiator(router)
	product.Initiator(router)
	order.Initiator(router)
	technician.Initiator(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
