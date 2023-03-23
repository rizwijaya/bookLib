package controllers

import (
	bookUseCase "bookLib/usecases"
)

type BookController struct {
	BookUseCase *bookUseCase.BookUseCase
}

func NewBookController() *BookController {
	bu := bookUseCase.NewBookUseCase()
	return &BookController{
		BookUseCase: bu,
	}
}
