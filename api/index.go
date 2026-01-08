package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/controllers"
)

var app *gin.Engine

func init() {
	app = gin.New()
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})
	controllers.InitController(app)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
