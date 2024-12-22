package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "net/http"
    "time"
)

func main() {
    // Create a new Gin router
    router := gin.Default()

	// Set up CORS middleware with custom configuration
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Allow specific origins
        AllowMethods:     []string{"GET", "POST"},                                // Allow specific HTTP methods
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},    // Allow specific headers
        ExposeHeaders:    []string{"Content-Length"},                             // Expose headers to client
        AllowCredentials: true,                                                   // Allow credentials like cookies
        MaxAge:           12 * time.Hour,                                         // Cache the preflight request for 12 hours
    }))

    // Health check endpoint
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "UP"})
    })

    // Start the server on port 9000
    if err := router.Run(":9000"); err != nil {
        panic(err)
    }
}
