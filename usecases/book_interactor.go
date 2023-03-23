package usecases

import (
	"bookLib/domain"
)

func (bu *BookUseCase) AllBooks() []domain.Book {
	return books
}
