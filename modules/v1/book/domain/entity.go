package domain

import "time"

type Book struct {
	ID         int       `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	Name_book  string    `json:"name_book" gorm:"column:name_book;type:varchar(100);not null"`
	Author     string    `json:"author" gorm:"column:author;type:varchar(100);not null"`
	Created_at time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	Updated_at time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
}

type Books []Book
