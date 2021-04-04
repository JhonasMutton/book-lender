package model

import "time"

type User struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Name          string     `json:"name" gorm:"size:100"`
	Email         string     `json:"email" gorm:"unique,size:150"`
	CreatedAt     time.Time  `json:"created_at"`
	Collection    []Book     `json:"collection,omitempty" gorm:"foreignKey:OwnerId"`
	LentBooks     []LoanBook `json:"lent_books,omitempty" gorm:"foreignKey:FromUser"`
	BorrowedBooks []LoanBook `json:"borrowed_books,omitempty" gorm:"foreignKey:ToUser"`
}

type UserDto struct {
	Name  string `json:"name" validate:"required,max=100"`
	Email string `json:"email" gorm:"unique" validate:"required,email,max=150"`
}

func (dto UserDto) ToModel() User {
	return User{
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type Users []User
