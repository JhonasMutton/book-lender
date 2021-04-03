package model

import (
	"strconv"
	"time"
)

type Book struct {
	ID          uint       `json:"id"  gorm:"primaryKey"`
	Title       string     `json:"title"`
	Pages       uint       `json:"pages"`
	CreatedAt   time.Time  `json:"created_at"`
	OwnerId     uint       `json:"-"`
	BookHistory []LoanBook `json:"book_history,omitempty" gorm:"foreignKey:BookID"`
}

type BookDTO struct {
	Title        string `json:"title"`
	Pages        string `json:"pages"`
	LoggedUserId uint   `json:"logged_user_id"`
}

func (dto BookDTO) ToModel() Book {
	pages, _ := strconv.ParseUint(dto.Pages, 10, 32) //TODO Verificar se validator verifica isso
	return Book{
		Title:   dto.Title,
		Pages:   uint(pages),
		OwnerId: dto.LoggedUserId,
	}
}
