package usecases

import (
	"bookLib/modules/v1/book/domain"
	"time"
)

func (bu *BookUseCase) AllBooks() (domain.Books, error) {
	return bu.repoBook.AllBooks()
}

func (bu *BookUseCase) GetBookByID(id string) (domain.Book, error) {
	return bu.repoBook.GetBookByID(id)
}

func (bu *BookUseCase) CreateBook(book domain.InsertBook) (domain.Book, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return domain.Book{}, err
	}
	now := time.Now().In(location)
	books := domain.Book{
		Title:      book.Title,
		Author:     book.Author,
		Updated_at: now,
		Created_at: now,
	}

	return bu.repoBook.CreateBook(books)
}

func (bu *BookUseCase) UpdateBook(id string, input domain.UpdateBook) (domain.Book, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return domain.Book{}, err
	}
	now := time.Now().In(location)
	//id_book, err := strconv.Atoi(id)
	if err != nil {
		return domain.Book{}, err
	}
	book := domain.Book{
		Title:      input.Title,
		Author:     input.Author,
		Updated_at: now,
	}
	return bu.repoBook.UpdateBook(id, book)
}

func (bu *BookUseCase) DeleteBook(id string) error {
	return bu.repoBook.DeleteBook(id)
}
