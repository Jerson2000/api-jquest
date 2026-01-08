package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/enums"
	"github.com/jerson2000/jquest/responses"
	"github.com/jerson2000/jquest/utils"
)

func RBACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role").(enums.Role)

		path := c.Request.URL.Path
		method := c.Request.Method

		if config.CasbinEnforcer == nil {
			res := responses.Failure[any](http.StatusInternalServerError, "Something went wrong")
			c.AbortWithStatusJSON(utils.ToHTTPStatus(res.Status), res)
			return
		}

		ok, err := config.CasbinEnforcer.Enforce(string(role), path, method)
		log.Printf("DEBUG: Enforce: sub=%q obj=%q act=%q allowed=%v err=%v",
			role, path, method, ok, err)
		if err != nil {
			res := responses.Failure[any](http.StatusInternalServerError, "Something went wrong")
			c.AbortWithStatusJSON(utils.ToHTTPStatus(res.Status), res)
			return
		}

		if !ok {
			res := responses.Failure[any](http.StatusForbidden, "You don't have permission to access this resource!")
			c.AbortWithStatusJSON(utils.ToHTTPStatus(res.Status), res)
			return
		}

		c.Next()
	}
}
