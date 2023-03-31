package controllers

import (
	bookRepository "bookLib/modules/v1/book/interfaces/repository"
	bookUseCase "bookLib/modules/v1/book/usecases"
	"database/sql"
)

type BookController struct {
	BookUseCase *bookUseCase.BookUseCase
}

func NewBookController(db *sql.DB) *BookController {
	repo := bookRepository.NewBookRepository(db)
	bu := bookUseCase.NewBookUseCase(repo)
	return &BookController{
		BookUseCase: bu,
	}
}
