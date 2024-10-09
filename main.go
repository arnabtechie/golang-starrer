package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arnabtechie/golang-starrer/database"
	"github.com/arnabtechie/golang-starrer/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	defer database.CloseDB()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}

	r.Use(cors.New(corsConfig))

	database.InitDB(os.Getenv("CONN_STR"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "alive",
		})
	})

	routes.SetupRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Page not found",
			"message": "The resource you are looking for does not exist.",
		})
	})

	r.Run(":" + os.Getenv("PORT"))
}
