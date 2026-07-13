package handler

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/controllers"
)

var (
	app  *gin.Engine
	once sync.Once
)

func initApp() {
	once.Do(func() {
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
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	initApp()
	app.ServeHTTP(w, r)
}
