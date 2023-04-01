package controllers

import (
	bookRepository "bookLib/modules/v1/book/interfaces/repositories"
	bookUseCase "bookLib/modules/v1/book/usecases"

	"gorm.io/gorm"
)

type BookController struct {
	BookUseCase *bookUseCase.BookUseCase
}

func NewBookController(db *gorm.DB) *BookController {
	repo := bookRepository.NewBookRepository(db)
	bu := bookUseCase.NewBookUseCase(repo)
	return &BookController{
		BookUseCase: bu,
	}
}
