package controllers

import (
	api "bookLib/pkg/api_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (bc *BookController) GetBooks(c *gin.Context) {
	book := bc.BookUseCase.AllBooks()
	response := api.APIResponse("Success to get all books!", http.StatusOK, "success", book)
	c.JSON(http.StatusOK, response)
}
