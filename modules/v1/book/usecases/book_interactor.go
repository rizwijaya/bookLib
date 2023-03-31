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

func (bu *BookUseCase) CreateBook(book domain.Book) (domain.Book, error) {
	return bu.repoBook.CreateBook(book)
}

func (bu *BookUseCase) UpdateBook(id string, book domain.Book) (domain.Book, error) {
	return bu.repoBook.UpdateBook(id, book)
}

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
