package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jerson2000/jquest/responses"
)

func ToHTTPStatus(status int) int {
	switch status {
	case http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNoContent,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusConflict,
		http.StatusUnprocessableEntity,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable:
		return status
	default:
		return http.StatusInternalServerError
	}
}

func ValidationShouldBind[T any](status int, dto T, trans ut.Translator, c *gin.Context) bool {
	if err := c.ShouldBind(&dto); err != nil {
		var verrs validator.ValidationErrors

		if errors.As(err, &verrs) {
			out := make(map[string]string)
			for _, fe := range verrs {
				out[fe.Field()] = fe.Translate(trans)
			}
			res := responses.Failures[any](http.StatusBadRequest, out)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return false
		}

		res := responses.Failure[any](status, err.Error())
		c.AbortWithStatusJSON(ToHTTPStatus(status), res)
		return false
	}
	return true
}
