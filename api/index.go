package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/internal/controllers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	controllers.InitController(router)
	router.ServeHTTP(w, r)
}
