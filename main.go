package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	fmt.Println("Server berjalan pada http://localhost:8080")
	select {}
}
