package usecases

import (
	"bookLib/modules/v1/book/domain"
	bookRepository "bookLib/modules/v1/book/interfaces/repositories"
)

type BookAdapter interface {
	AllBooks() ([]domain.Book, error)
	GetBookByID(id string) (domain.Book, error)
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id string, book domain.Book) (domain.Book, error)
	DeleteBook(id string) error
}

type BookUseCase struct {
	repoBook *bookRepository.Repository
}

func NewBookUseCase(repoBook *bookRepository.Repository) *BookUseCase {
	return &BookUseCase{repoBook}
}
