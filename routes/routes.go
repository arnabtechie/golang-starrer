package routes

import (
	"net/http"

	"github.com/arnabtechie/golang-starrer/controllers/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	user := api.Group("/user")

	user.POST("/register", func(c *gin.Context) {
		var data auth.User

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := auth.Register(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "data": response})
	})
}
