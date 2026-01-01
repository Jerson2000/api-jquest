package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Exported handler for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Use a different name for the Gin router (e.g., router or engine)
	router := gin.Default()

	// Define your routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin on Vercel! (Working version)",
		})
	})

	router.GET("/api/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "World")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	router.POST("/api/echo", func(c *gin.Context) {
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, body)
	})

	// Correct: pass the original *http.Request as second argument
	router.ServeHTTP(w, r)
}
