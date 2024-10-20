package main

import (
	"fmt"
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
	user.Initiator(router)
	product.Initiator(router)
	order.Initiator(router)
	technician.Initiator(router)
	router.SetTrustedProxies(nil)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	fmt.Println("Server berjalan pada http://localhost:8080")
	select {}

}
