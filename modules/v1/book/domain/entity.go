package domain

type Book struct {
	ID     int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	Title  string `json:"title" binding:"required,min=3" gorm:"column:title;type:varchar(100);not null"`
	Author string `json:"author" binding:"required,min=3" gorm:"column:author;type:varchar(100);not null"`
	Desc   string `json:"desc" binding:"required,min=8" gorm:"column:description;type:varchar(400);not null"`
}

type Books []Book
