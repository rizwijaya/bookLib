package domain

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" binding:"required,min=3"`
	Author string `json:"author" binding:"required,min=3"`
	Desc   string `json:"desc" binding:"required,min=8"`
}

type Books []Book
