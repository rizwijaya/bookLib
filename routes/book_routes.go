package routes

import (
	bookControllerV1 "bookLib/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) *gin.Engine {
	bookControllerV1 := bookControllerV1.NewBookController()
	api := router.Group("/books")
	api.GET("/", bookControllerV1.GetBooks)
	api.GET("/:id", bookControllerV1.GetBookByID)
	api.POST("/", bookControllerV1.AddBook)
	api.PUT("/:id", bookControllerV1.UpdateBook)

	return router
}
