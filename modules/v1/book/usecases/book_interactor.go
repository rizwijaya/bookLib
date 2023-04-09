package usecases

import (
	"bookLib/modules/v1/book/domain"
	"bookLib/pkg/times"
)

func (bu *BookUseCase) AllBooks() (domain.Books, error) {
	return bu.repoBook.AllBooks()
}

func (bu *BookUseCase) GetBookByID(id string) (domain.Book, error) {
	return bu.repoBook.GetBookByID(id)
}

func (bu *BookUseCase) CreateBook(book domain.InsertBook) (domain.Book, error) {
	now := times.Now("Asia/Jakarta")
	books := domain.Book{
		Name_book:  book.Name_book,
		Author:     book.Author,
		Updated_at: now,
		Created_at: now,
	}

	return bu.repoBook.CreateBook(books)
}

func (bu *BookUseCase) UpdateBook(id string, input domain.UpdateBook) (domain.Book, error) {
	now := times.Now("Asia/Jakarta")
	book := domain.Book{
		Name_book:  input.Name_book,
		Author:     input.Author,
		Updated_at: now,
	}
	//Check Book Exist
	_, err := bu.repoBook.GetBookByID(id)
	if err != nil {
		return domain.Book{}, err
	}
	return bu.repoBook.UpdateBook(id, book)
}

func (bu *BookUseCase) DeleteBook(id string) error {
	//Check Book Exist
	_, err := bu.repoBook.GetBookByID(id)
	if err != nil {
		return err
	}
	return bu.repoBook.DeleteBook(id)
}
