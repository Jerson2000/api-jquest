package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/controllers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin on Vercel! (Working version)",
		})
	})
	controllers.InitController(router)
	router.ServeHTTP(w, r)
}
