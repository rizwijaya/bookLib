package domain

import "time"

type Book struct {
	ID         int       `gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	Title      string    `gorm:"column:title;type:varchar(100);not null"`
	Author     string    `gorm:"column:author;type:varchar(100);not null"`
	Updated_at time.Time `gorm:"column:updated_at;type:timestamp;not null"`
	Created_at time.Time `gorm:"column:created_at;type:timestamp;not null"`
}

type Books []Book
