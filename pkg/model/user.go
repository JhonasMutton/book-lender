package model

import "time"

type BasicUser struct {
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	BasicUser
	CreatedAt     time.Time   `json:"created_at"`
	Collection    []Book      `json:"collection,omitempty" gorm:"foreignKey:OwnerId"`
	LentBooks     []LoanBooks `json:"lent_books,omitempty" gorm:"foreignKey:FromUser"`
	BorrowedBooks []LoanBooks `json:"borrowed_books,omitempty" gorm:"foreignKey:ToUser"`
}

type Users []User

type Book struct {
	ID          uint        `json:"id"  gorm:"primaryKey"`
	Title       string      `json:"title"`
	Pages       uint        `json:"pages"`
	CreatedAt   time.Time   `json:"created_at"`
	OwnerId     uint        `json:"owner_id"`
	BookHistory []LoanBooks `json:"borrowed_books" gorm:"foreignKey:BookID"`
}

type LoanBooks struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	BookID     uint      `json:"book_id"`
	FromUser   uint      `json:"from_user"`
	ToUser     uint      `json:"to_user"`
	LentAt     time.Time `json:"lent_at"`
	ReturnedAt time.Time `json:"returned_at"`
}
