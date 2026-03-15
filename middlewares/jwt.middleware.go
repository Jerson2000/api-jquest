package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/models"
	"github.com/jerson2000/jquest/responses"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			res := responses.Failure[any](http.StatusUnauthorized, "You don't have permission to access this resource!")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if c.FullPath() == "/api/auth/refresh" {
			if !isJwtFormat(tokenString) {
				res := responses.Failure[any](http.StatusUnauthorized, "invalid token format")
				c.JSON(http.StatusUnauthorized, res)
				c.Abort()
				return
			}
			c.Next()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})

		if err != nil || !token.Valid {
			res := responses.Failure[any](http.StatusUnauthorized, "invalid token")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.JWTClaims); ok {
			c.Set("id", claims.Id)
			c.Set("name", claims.Name)
			c.Set("role", claims.Role)
			c.Next()
		} else {
			res := responses.Failure[any](http.StatusUnauthorized, "invalid token claims")
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
		}
	}
}

func isJwtFormat(token string) bool {
	parts := strings.Split(token, ".")
	return len(parts) == 3
}
