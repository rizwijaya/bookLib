package error

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Form struct {
	Field   string
	Message string
}

func PageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(404, "error_404.html", nil)
	}
}

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(440, gin.H{"status": "error 440", "message": "Method Not Allowed"})
	}
}

func FormValidationError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required!"
	case "email":
		return fe.Field() + " must be a valid email address!"
	case "min":
		return fe.Field() + " minimum " + fe.Param() + " characters!"
	case "max":
		return fe.Field() + " maximum " + fe.Param() + " characters!"
	case "alphanum":
		return fe.Field() + " must be alphanumeric!"
	case "numeric":
		return fe.Field() + " must be numeric!"
	case "eqfield":
		return fe.Field() + " must be equal to " + fe.Param() + "!"
	case "alphanumunicode":
		return fe.Field() + " must be alphanumeric unicode!"
	default:
		return fe.Field() + " is invalid!"
	}
}
