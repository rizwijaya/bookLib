package usecases

import (
	"bookLib/modules/v1/book/domain"
	bookRepository "bookLib/modules/v1/book/interfaces/repository"
)

type BookAdapter interface {
	AllBooks() []domain.Book
	GetBookByID(id string) domain.Book
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id string, book domain.Book) (domain.Book, error)
	DeleteBook(id string) error
}

type BookUseCase struct {
	repoBook *bookRepository.Repository
}

func NewBookUseCase(repoBook *bookRepository.Repository) *BookUseCase {
	return &BookUseCase{repoBook}
}

// var books = domain.Books{
// 	{
// 		ID:     1,
// 		Title:  "The Hobbit",
// 		Author: "J.R.R. Tolkien",
// 		Desc:   "The Hobbit is a children's fantasy novel by English author J. R. R. Tolkien. It was published on 21 September 1937 to wide critical acclaim, being nominated for the Carnegie Medal and awarded a prize from the New York Herald Tribune for best juvenile fiction. The book remains popular and is recognized as a classic in children's literature.",
// 	},
// 	{
// 		ID:     2,
// 		Title:  "The Lord of the Rings",
// 		Author: "J.R.R. Tolkien",
// 		Desc:   "The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien. The story began as a sequel to Tolkien's 1937 fantasy novel The Hobbit, but eventually developed into a much larger work. Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling novels ever written, with over 150 million copies sold.",
// 	},
// 	{
// 		ID:     3,
// 		Title:  "The Chronicles of Narnia",
// 		Author: "C.S. Lewis",
// 		Desc:   "The Chronicles of Narnia is a series of seven fantasy novels by C. S. Lewis. It is considered a classic of children's literature and is the author's best-known work, having sold over 100 million copies in 47 languages. The series was written by Lewis between 1949 and 1954, although it was not published until 1950. The books are set in the fictional realm of Narnia, a fantasy world of magic, mythical beasts, and talking animals.",
// 	},
// 	// {
// 	// 	ID:     4,
// 	// 	Title:  "The Lion, the Witch and the Wardrobe",
// 	// 	Author: "C.S. Lewis",
// 	// 	Desc:   "The Lion, the Witch and the Wardrobe is a fantasy novel for children by C. S. Lewis, published by Geoffrey Bles in 1950. It is the first published and best known of seven novels in The Chronicles of Narnia (1950–1956). It is considered a classic of children's literature and is the author's best-known work, having sold over 100 million copies in 47 languages.",
// 	// },
// }
