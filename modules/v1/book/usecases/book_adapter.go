package usecases

import (
	"bookLib/modules/v1/book/domain"
	bookRepository "bookLib/modules/v1/book/interfaces/repositories"
)

type BookAdapter interface {
	AllBooks() (domain.Books, error)
	GetBookByID(id string) (domain.Book, error)
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id string, input domain.UpdateBook) (domain.Book, error)
	DeleteBook(id string) error
}

type BookUseCase struct {
	repoBook bookRepository.RepositoryPresenter
}

func NewBookUseCase(repoBook bookRepository.RepositoryPresenter) *BookUseCase {
	return &BookUseCase{repoBook}
}
