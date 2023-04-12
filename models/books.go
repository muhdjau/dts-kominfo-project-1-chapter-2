package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NameBook  string    `json:"name_book" gorm:"not null;type:varchar(50)"`
	Author    string    `json:"author" gorm:"not null;type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Books) TableName() string {
	return "books"
}

func (b *Books) BeforeCreate(tx *gorm.DB) (err error) {

	if len(b.NameBook) <= 2 {
		err = errors.New("Book's name is too short")
	}

	if len(b.Author) <= 2 {
		err = errors.New("Author's name is too short")
	}

	return
}
