package routes

import (
	bookControllerV1 "bookLib/modules/v1/book/interfaces/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, db *sql.DB) *gin.Engine {
	bookControllerV1 := bookControllerV1.NewBookController(db)
	api := router.Group("/books")
	api.GET("/", bookControllerV1.GetBooks)
	api.GET("/:id", bookControllerV1.GetBookByID)
	// api.POST("/", bookControllerV1.AddBook)
	// api.PUT("/:id", bookControllerV1.UpdateBook)
	// api.DELETE("/:id", bookControllerV1.DeleteBook)

	return router
}
