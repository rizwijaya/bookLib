package usecases

import (
	"bookLib/modules/v1/book/domain"
)

func (bu *BookUseCase) AllBooks() ([]domain.Book, error) {
	return bu.repoBook.AllBooks()
}

func (bu *BookUseCase) GetBookByID(id string) (domain.Book, error) {
	return bu.repoBook.GetBookByID(id)
}

// func (bu *BookUseCase) CreateBook(book domain.Book) (domain.Book, error) {
// 	book.ID = len(books) + 1
// 	books = append(books, book)
// 	return book, nil
// }

// func (bu *BookUseCase) UpdateBook(id string, book domain.Book) (domain.Book, error) {
// 	var err error
// 	book.ID, err = strconv.Atoi(id)
// 	if err != nil {
// 		return domain.Book{}, err
// 	}
// 	for i, b := range books {
// 		if b.ID == book.ID {
// 			books[i] = book
// 			return book, nil
// 		}
// 	}
// 	return domain.Book{}, errors.New("book not found")
// }

// func (bu *BookUseCase) DeleteBook(id string) error {
// 	id_book, _ := strconv.Atoi(id)
// 	for i, b := range books {
// 		if b.ID == id_book {
// 			books = append(books[:i], books[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return errors.New("book not found")
// }
