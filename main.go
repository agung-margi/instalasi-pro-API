package main

import (
	"fmt"
	"instalasi-pro/database"
	"instalasi-pro/modules/order"
	"instalasi-pro/modules/product"
	"instalasi-pro/modules/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connection()
	database.DB.AutoMigrate(&user.User{})
	database.DB.AutoMigrate(&product.Product{})
	database.DB.AutoMigrate(&order.Order{})

	router := gin.Default()
	user.Initiator(router)
	product.Initiator(router)
	order.Initiator(router)
	router.SetTrustedProxies(nil)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	fmt.Println("Server berjalan pada http://localhost:8080")
	select {}

}
