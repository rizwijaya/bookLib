package repository

import (
	"bookLib/modules/v1/book/domain"
	"errors"
)

func (r *Repository) AllBooks() ([]domain.Book, error) {
	var books []domain.Book
	rows, err := r.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *Repository) GetBookByID(id string) (domain.Book, error) {
	var book domain.Book
	err := r.db.QueryRow("SELECT * FROM books WHERE id = "+id).Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *Repository) CreateBook(book domain.Book) (domain.Book, error) {
	_, err := r.db.Exec("INSERT INTO books (title, author, description) VALUES ($1, $2, $3)", book.Title, book.Author, book.Desc)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *Repository) UpdateBook(id string, book domain.Book) (domain.Book, error) {
	res, err := r.db.Exec("UPDATE books SET title = $1, author = $2, description = $3 WHERE id = "+id, book.Title, book.Author, book.Desc)
	if err != nil {
		return domain.Book{}, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return domain.Book{}, err
	}
	if rowsAffected == 0 {
		return domain.Book{}, errors.New("book not found")
	}
	return book, nil
}

func (r *Repository) DeleteBook(id string) error {
	res, err := r.db.Exec("DELETE FROM books WHERE id = " + id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}
