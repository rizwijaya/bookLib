package controllers

import (
	"bookLib/modules/v1/book/domain"
	error "bookLib/pkg/http-error"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (bc *BookController) GetBooks(c *gin.Context) {
	book, err := bc.BookUseCase.AllBooks()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Failed to Get Book!")
		return
	}
	if len(book) == 0 {
		c.JSON(http.StatusOK, "Book Not Found!")
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := bc.BookUseCase.GetBookByID(id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Failed to Get Book!")
		return
	}
	if book == (domain.Book{}) {
		c.JSON(http.StatusOK, "Book Not Found!")
		return
	}
	c.JSON(http.StatusOK, book)
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
			log.Println(res)
			c.JSON(http.StatusBadRequest, res)
			return
		}
	}

	book, err := bc.BookUseCase.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to Add New Book!")
		return
	}

	c.JSON(http.StatusOK, "Created")
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
			log.Println(res)
			c.JSON(http.StatusBadRequest, res)
			return
		}
	}

	book, err := bc.BookUseCase.UpdateBook(id, book)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Failed to Update Book!")
		return
	}

	c.JSON(http.StatusOK, "Updated")
}

// func (bc *BookController) DeleteBook(c *gin.Context) {
// 	id := c.Param("id")
// 	err := bc.BookUseCase.DeleteBook(id)
// 	if err != nil {
// 		response := api.APIResponse("Failed to Delete Book!", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	response := api.APIResponse("Success to Delete Book!", http.StatusOK, "success", nil)
// 	c.JSON(http.StatusOK, response)
// }
