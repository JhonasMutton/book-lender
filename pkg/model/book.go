package model

import (
	"strconv"
	"time"
)

type Book struct {
	ID          uint       `json:"id"  gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"size:200"`
	Pages       uint       `json:"pages"`
	CreatedAt   time.Time  `json:"created_at"`
	Owner       uint       `json:"-"`
	BookHistory []LoanBook `json:"book_history,omitempty" gorm:"foreignKey:Book"`
}

type BookDTO struct {
	Title        string `json:"title" validate:"required,max=200"`
	Pages        string `json:"pages" validate:"required,number"`
	LoggedUserId uint   `json:"logged_user_id" validate:"required"`
}

func (dto BookDTO) ToModel() Book {
	pages, _ := strconv.ParseUint(dto.Pages, 10, 32) //TODO Verificar se validator verifica isso
	return Book{
		Title: dto.Title,
		Pages: uint(pages),
		Owner: dto.LoggedUserId,
	}
}
