package repository

import "bookLib/modules/v1/book/domain"

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
