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
		return fe.Field() + " wajib diisi!"
	case "email":
		return fe.Field() + " harus berupa email!"
	case "min":
		return fe.Field() + " minimal " + fe.Param() + " karakter!"
	case "max":
		return fe.Field() + " maksimal " + fe.Param() + " karakter!"
	case "alphanum":
		return fe.Field() + " hanya boleh berisi huruf dan angka!"
	case "numeric":
		return fe.Field() + " hanya boleh berisi angka!"
	case "eqfield":
		return fe.Field() + " harus sama dengan " + fe.Param() + "!"
	case "alphanumunicode":
		return fe.Field() + " harus berisi karakter, huruf dan angka!"
	default:
		return fe.Field() + " tidak valid!"
	}
}
