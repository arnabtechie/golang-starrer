package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	type User struct {
		Name    string
		Mob     string
		address string
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    User{Name: "arnab", Mob: "7001047915", address: "Purulia"},
		})
	})
	r.Run(":" + os.Getenv("PORT"))
}
