package model

import "time"

type User struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Name          string     `json:"name"`
	Email         string     `json:"email" gorm:"unique"`
	CreatedAt     time.Time  `json:"created_at"`
	Collection    []Book     `json:"collection,omitempty" gorm:"foreignKey:OwnerId"`
	LentBooks     []LoanBook `json:"lent_books,omitempty" gorm:"foreignKey:FromUser"`
	BorrowedBooks []LoanBook `json:"borrowed_books,omitempty" gorm:"foreignKey:ToUser"`
}

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

func (dto UserDto) ToModel() User {
	return User{
		Name:  dto.Name,
		Email: dto.Email,
	}
}

type Users []User

