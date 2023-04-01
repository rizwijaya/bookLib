package routes

import (
	_ "bookLib/docs"
	bookControllerV1 "bookLib/modules/v1/book/interfaces/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	bookControllerV1 := bookControllerV1.NewBookController(db)
	api := router.Group("/books")
	{
		api.GET("", bookControllerV1.GetBooks)
		api.GET("/:id", bookControllerV1.GetBookByID)
		api.POST("/", bookControllerV1.AddBook)
		api.PUT("/:id", bookControllerV1.UpdateBook)
		api.DELETE("/:id", bookControllerV1.DeleteBook)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
