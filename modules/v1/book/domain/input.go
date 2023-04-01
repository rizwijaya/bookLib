package domain

type UpdateBook struct {
	Name_book string `json:"name_book" binding:"omitempty,min=3" example:"The Lord of the Rings"`
	Author    string `json:"author" binding:"omitempty,min=3" example:"J.R.R. Tolkien"`
}

type InsertBook struct {
	Name_book string `json:"name_book" binding:"required,min=3" example:"The Lord of the Rings"`
	Author    string `json:"author" binding:"required,min=3" example:"J.R.R. Tolkien"`
}
