package repository

import (
	"bookLib/modules/v1/book/domain"

	"gorm.io/gorm"
)

type RepositoryPresenter interface {
	AllBooks() ([]domain.Book, error)
	GetBookByID(id string) (domain.Book, error)
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id string, book domain.Book) (domain.Book, error)
	DeleteBook(id string) error
}

type Repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
