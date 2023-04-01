package repository

import (
	"bookLib/modules/v1/book/domain"

	"gorm.io/gorm/clause"
)

func (r *Repository) AllBooks() (domain.Books, error) {
	var books domain.Books
	err := r.db.Find(&books).Error
	return books, err
}

func (r *Repository) GetBookByID(id string) (domain.Book, error) {
	var book domain.Book
	err := r.db.Where("id = " + id).First(&book).Error
	return book, err
}

func (r *Repository) CreateBook(book domain.Book) (domain.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *Repository) UpdateBook(id string, book domain.Book) (domain.Book, error) {
	res := r.db.Model(&book).Clauses(clause.Returning{}).Where("id = " + id).Updates(book)
	return book, res.Error
}

func (r *Repository) DeleteBook(id string) error {
	var book domain.Book
	return r.db.Delete(&book, id).Error
}
