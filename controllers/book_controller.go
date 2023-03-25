package controllers

import (
	"bookLib/domain"
	api "bookLib/pkg/api_response"
	error "bookLib/pkg/http-error"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (bc *BookController) GetBooks(c *gin.Context) {
	book := bc.BookUseCase.AllBooks()
	if len(book) == 0 {
		response := api.APIResponse("Book Not Found!", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusOK, response)
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

func (bc *BookController) AddBook(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]error.Form, len(verr))
			for i, fe := range verr {
				res[i] = error.Form{
					Field:   fe.Field(),
					Message: error.FormValidationError(fe),
				}
			}
			response := api.APIResponse("Failed to Add New Book!", http.StatusBadRequest, "error", res)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	book, err := bc.BookUseCase.CreateBook(book)
	if err != nil {
		response := api.APIResponse("Failed to Add New Book!", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := api.APIResponse("Success to Add New Book!", http.StatusOK, "success", book)
	c.JSON(http.StatusOK, response)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	var book domain.Book
	id := c.Param("id")
	if err := c.ShouldBindJSON(&book); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]error.Form, len(verr))
			for i, fe := range verr {
				res[i] = error.Form{
					Field:   fe.Field(),
					Message: error.FormValidationError(fe),
				}
			}
			response := api.APIResponse("Failed to Update Book!", http.StatusBadRequest, "error", res)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	book, err := bc.BookUseCase.UpdateBook(id, book)
	if err != nil {
		response := api.APIResponse("Failed to Update Book!", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := api.APIResponse("Success to Update Book!", http.StatusOK, "success", book)
	c.JSON(http.StatusOK, response)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := bc.BookUseCase.DeleteBook(id)
	if err != nil {
		response := api.APIResponse("Failed to Delete Book!", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := api.APIResponse("Success to Delete Book!", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
