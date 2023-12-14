package main

import (
	"latihan1-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Setup route untuk modul User
	r.GET("/users", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)

	// GIN server
	r.Run(":8080")
}
