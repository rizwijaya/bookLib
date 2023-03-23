package controllers

import (
	"github.com/gin-gonic/gin"
)

func (bc *BookController) GetBooks(c *gin.Context) {
	book := bc.BookUseCase.AllBooks()
	c.JSON(200, book)
}
