package repository

import (
	"bookLib/modules/v1/book/domain"
)

func (r *Repository) AllBooks() ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *Repository) GetBookByID(id string) (domain.Book, error) {
	var book domain.Book
	err := r.db.Where("id = " + id).First(&book).Error
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *Repository) CreateBook(book domain.Book) (domain.Book, error) {
	result := r.db.Create(&book)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *Repository) UpdateBook(id string, book domain.Book) (domain.Book, error) {
	err := r.db.Model(&book).Where("id = " + id).Updates(book).Error
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *Repository) DeleteBook(id string) error {
	var book domain.Book
	err := r.db.Where("id = " + id).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
