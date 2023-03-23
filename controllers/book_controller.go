package controllers

import (
	api "bookLib/pkg/api_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BookController) GetBooks(c *gin.Context) {
	book := bc.BookUseCase.AllBooks()
	if book == nil {
		response := api.APIResponse("Book Not Found!", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := api.APIResponse("Success to Get All Books!", http.StatusOK, "success", book)
	c.JSON(http.StatusOK, response)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book := bc.BookUseCase.GetBookByID(id)
	if len(book.Title) == 0 {
		response := api.APIResponse("Book Not Found!", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := api.APIResponse("Success to Get Book!", http.StatusOK, "success", book)
	c.JSON(http.StatusOK, response)
}
