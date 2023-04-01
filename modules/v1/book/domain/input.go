package domain

type UpdateBook struct {
	Title  string `json:"title" binding:"omitempty,min=3" example:"The Lord of the Rings"`
	Author string `json:"author" binding:"omitempty,min=3" example:"J.R.R. Tolkien"`
}

type InsertBook struct {
	Title  string `json:"title" binding:"required,min=3" example:"The Lord of the Rings"`
	Author string `json:"author" binding:"required,min=3" example:"J.R.R. Tolkien"`
}
