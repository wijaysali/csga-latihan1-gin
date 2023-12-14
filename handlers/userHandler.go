package handlers

import (
	"latihan1-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
}

// func GetUser(c *gin.Context) {
// 	// kembalikan list user
// 	c.JSON(http.StatusOK, users)
// }

func GetUser(c *gin.Context) {
	if len(users) == 0 {
		c.JSON(http.StatusOK, []string{})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	//parse JSON dari req body
	// if err := c.ShouldBindJSON(&newUser); err != nil{
	// 	c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
	// }

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users = append(users, newUser)

	//return user baru
	c.JSON(http.StatusCreated, newUser)
}
