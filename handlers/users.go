package handlers

import (
	"log"
	"net/http"
	"backend-3/db"
	"backend-3/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Initialize the database connection
	database := db.InitializeDB()

	// Fetch users from the database
	var users []models.User

	if err := database.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Return the users as JSON response
	c.JSON(http.StatusOK, users)
	log.Println(users)
}