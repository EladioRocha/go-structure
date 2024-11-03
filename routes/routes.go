package routes

import (
	"backend-3/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", handlers.GetUsers)
}