package usecases

import (
	"bookLib/domain"
	"errors"
	"strconv"
)

func (bu *BookUseCase) AllBooks() []domain.Book {
	return books
}

func (bu *BookUseCase) GetBookByID(id string) domain.Book {
	id_book, _ := strconv.Atoi(id)
	for _, book := range books {
		if book.ID == id_book {
			return book
		}
	}
	return domain.Book{}
}

func (bu *BookUseCase) CreateBook(book domain.Book) (domain.Book, error) {
	book.ID = len(books) + 1
	books = append(books, book)
	return book, nil
}

func (bu *BookUseCase) UpdateBook(id string, book domain.Book) (domain.Book, error) {
	id_book, _ := strconv.Atoi(id)
	for i, b := range books {
		if b.ID == id_book {
			books[i] = book
			return book, nil
		}
	}
	return domain.Book{}, errors.New("book not found")
}

func (bu *BookUseCase) DeleteBook(id string) error {
	id_book, _ := strconv.Atoi(id)
	for i, b := range books {
		if b.ID == id_book {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
