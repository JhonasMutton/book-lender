package model

import "time"

type LoanBook struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Book       uint      `json:"book"`
	FromUser   uint      `json:"from_user"`
	ToUser     uint      `json:"to_user"`
	LentAt     time.Time `json:"lent_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `json:"Status" gorm:"type:ENUM('lent', 'returned')"`
}

//Status
const (
	StatusLent     = "lent"
	StatusReturned = "returned"
)

type LendBookDTO struct {
	Book       uint `json:"book_id" validate:"required"`
	LoggedUser uint `json:"logged_user_id" validate:"required"`
	ToUser     uint `json:"to_user_id" validate:"required"`
}

func (l LendBookDTO) ToModel() LoanBook {
	return LoanBook{
		Book:       l.Book,
		FromUser:   l.LoggedUser,
		ToUser:     l.ToUser,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     StatusLent,
	}
}

type ReturnBookDTO struct {
	BookID     uint `json:"book_id" validate:"required"`
	LoggedUser uint `json:"logged_user_id" validate:"required"`
}

func (l ReturnBookDTO) ToModel() LoanBook {
	return LoanBook{
		Book:   l.BookID,
		ToUser: l.LoggedUser,
		Status: StatusLent,
	}
}
