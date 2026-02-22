package utils

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// SetupTestRouter creates a gin engine for testing purposes
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

// CreateTestContext returns a fresh gin context and recorder
func CreateTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}
