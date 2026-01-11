package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"github.com/jerson2000/jquest/responses"
)

func CSRFMiddleware(authKey []byte, secure bool) gin.HandlerFunc {
	errorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		reason := csrf.FailureReason(r)
		log.Println("CSRFFailure:", reason)
		b, _ := json.Marshal(responses.Failure[any](
			http.StatusForbidden,
			"invalid token",
		))
		w.Write(b)
	})

	csrfMiddleware := csrf.Protect(
		authKey,
		csrf.CookieName("_forgery.anti"),
		csrf.Secure(secure),
		csrf.Path("/"),
		csrf.HttpOnly(true),
		csrf.ErrorHandler(errorHandler),
	)

	return func(c *gin.Context) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Set("csrf_token", csrf.Token(r))
			c.Next()
		})

		csrfMiddleware(handler).ServeHTTP(c.Writer, c.Request)

		if c.Writer.Status() == http.StatusForbidden {
			c.Abort()
		}
	}
}
