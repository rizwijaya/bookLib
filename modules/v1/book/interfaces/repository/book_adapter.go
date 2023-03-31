package repository

import "database/sql"

type RepositoryPresenter interface {
}

type Repository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *Repository {
	return &Repository{db}
}
