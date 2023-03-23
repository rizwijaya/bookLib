package usecases

import (
	"bookLib/domain"
	"strconv"
)

func (bu *BookUseCase) AllBooks() []domain.Book {
	return books
}

func (bu *BookUseCase) GetBookByID(id string) domain.Book {
	id_book, _ := strconv.Atoi(id)
	for _, book := range books {
		if book.ID == id_book {
			return book
		}
	}
	return domain.Book{}
}
