package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/controllers"
)

var app *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.InitConfig()
	app = gin.New()
	app.Use(gin.Recovery())
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
