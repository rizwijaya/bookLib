package repository

import (
	"bookLib/modules/v1/book/domain"
	"database/sql"
)

type RepositoryPresenter interface {
	AllBooks() ([]domain.Book, error)
	GetBookByID(id string) (domain.Book, error)
}

type Repository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *Repository {
	return &Repository{db}
}
