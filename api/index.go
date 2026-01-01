package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ginHandler(engine *gin.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		engine.ServeHTTP(w, r)
	}
}

func main() {
	// Create Gin engine
	r := gin.Default()

	// Sample routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin on Vercel!",
		})
	})

	r.GET("/api/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	r.POST("/api/echo", func(c *gin.Context) {
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, body)
	})

	// Vercel expects an exported Handler function with this signature
	http.HandleFunc("/", ginHandler(r))

	// Optional: Use port from Vercel env, fallback for local
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
