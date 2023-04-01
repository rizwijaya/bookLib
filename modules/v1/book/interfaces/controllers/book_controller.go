package controllers

import (
	"bookLib/modules/v1/book/domain"
	api "bookLib/pkg/api_response"
	error "bookLib/pkg/http-error"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//	@Summary		Get All Books
//	@Description	Get All Books
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int	true	"Limit of the book"
//	@Success		200		{object}	domain.Books
//	@Failure		404		{object}	api.Message
//	@Failure		500		{object}	api.Message
//	@Router			/books [get]
func (bc *BookController) GetBooks(c *gin.Context) {
	book, err := bc.BookUseCase.AllBooks()
	if err != nil {
		log.Println(err)
		res := api.SetMessage("Failed to Get Books!")
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if len(book) == 0 {
		res := api.SetMessage("Book Not Found!")
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, book)
}

// GetBookByID godoc
//	@Summary		Get Book By ID
//	@Description	Get Book By ID
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"ID of the book"
//	@Param			limit	path		int	true	"Limit of the book"
//	@Success		200		{object}	domain.Book
//	@Failure		404		{object}	api.Message
//	@Failure		500		{object}	api.Message
//	@Router			/books/{id} [get]
func (bc *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := bc.BookUseCase.GetBookByID(id)
	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			res := api.SetMessage("Book Not Found!")
			c.JSON(http.StatusNotFound, res)
			return
		}
		res := api.SetMessage("Failed to Get Book!")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, book)
}

// AddBook godoc
//	@Summary		Add Book
//	@Description	Add Book
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			title	query		string	true	"Title of the book"
//	@Param			author	query		string	true	"Author of the book"
//	@Success		200		{object}	domain.Book
//	@Failure		400		{object}	error.Form
//	@Failure		500		{object}	api.Message
//	@Router			/books [post]
func (bc *BookController) AddBook(c *gin.Context) {
	var input domain.InsertBook
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]error.Form, len(verr))
			for i, fe := range verr {
				res[i] = error.Form{
					Field:   fe.Field(),
					Message: error.FormValidationError(fe),
				}
			}
			c.JSON(http.StatusBadRequest, res)
			return
		}
		res := api.SetMessage("Failed to Create Book!")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	books, err := bc.BookUseCase.CreateBook(input)
	if err != nil {
		res := api.SetMessage("Failed to Create Book!")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, books)
}

// UpdateBook godoc
//	@Summary		Update Book
//	@Description	Update Book
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"ID of the book"
//	@Param			title	query		string	true	"Title of the book"
//	@Param			author	query		string	true	"Author of the book"
//	@Success		200		{object}	domain.Book
//	@Failure		400		{object}	error.Form
//	@Failure		404		{object}	api.Message
//	@Failure		500		{object}	api.Message
//	@Router			/books/{id} [put]
func (bc *BookController) UpdateBook(c *gin.Context) {
	var input domain.UpdateBook
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]error.Form, len(verr))
			for i, fe := range verr {
				res[i] = error.Form{
					Field:   fe.Field(),
					Message: error.FormValidationError(fe),
				}
			}
			c.JSON(http.StatusBadRequest, res)
			return
		}
		res := api.SetMessage("Failed to Create Book!")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if input.Name_book == "" && input.Author == "" {
		res := api.SetMessage("Name book and Author cannot be empty!")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	book, err := bc.BookUseCase.UpdateBook(id, input)
	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			res := api.SetMessage("Book Not Found!")
			c.JSON(http.StatusNotFound, res)
			return
		}
		res := api.SetMessage("Failed to Update Book!")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
//	@Summary		Delete Book
//	@Description	Delete Book
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID of the book"
//	@Success		200	{object}	api.Message
//	@Failure		400	{object}	api.Message
//	@Router			/books/{id} [delete]
func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := bc.BookUseCase.DeleteBook(id)
	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			res := api.SetMessage("Book Not Found!")
			c.JSON(http.StatusNotFound, res)
			return
		}
		res := api.SetMessage("Failed to Delete Book!")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := api.SetMessage("Book deleted successfully!")
	c.JSON(http.StatusOK, res)
}
