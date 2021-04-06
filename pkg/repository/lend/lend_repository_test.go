package lend

import (
	"fmt"
	"github.com/JhonasMutton/book-lender/pkg/database"
	"github.com/JhonasMutton/book-lender/pkg/model"
	testDatabase "github.com/JhonasMutton/book-lender/pkg/test/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

var (
	toUser     model.User
	fromUser   model.User
	book       model.Book
	db         *gorm.DB
	repository *Repository
)

func init() {
	testDatabase.InitMySqlContainer()
	setupTest()
}

func setupTest() {
	db = database.NewDatabaseConnection()
	repository = NewRepository(db)

	fromUser = model.User{
		Name:      "Bruce Wayne",
		Email:     "bruce.wayne@wayne.com",
		CreatedAt: time.Now(),
	}

	toUser = model.User{
		Name:      "Alfred Pennyworth",
		Email:     "alfred.pennyworth@wayne.com",
		CreatedAt: time.Now(),
	}

	result := db.Create(&fromUser)
	if err := result.Error; err != nil {
		panic("error to create fromUser" + err.Error())
	}

	result = db.Create(&toUser)
	if err := result.Error; err != nil {
		panic("error to create toUser" + err.Error())
	}

	book = model.Book{
		Title:     "Sherlock Holmes. O Cão dos Baskervilles",
		Pages:     176,
		CreatedAt: time.Now(),
		Owner:     fromUser.ID,
	}

	result = db.Create(&book)
	if err := result.Error; err != nil {
		panic("error to create book" + err.Error())
	}

}

func clearTable() {
	db.Exec("DELETE FROM loan_books")
}

func TestNewRepository(t *testing.T) {
	repository := NewRepository(db)

	assert.NotNil(t, repository)
	assert.NotNil(t, repository.db)
	assert.Equal(t, db, repository.db)
}

func TestRepository_Persist(t *testing.T) {
	defer clearTable()
	//given
	loanBook := model.LoanBook{
		Book:       book.ID,
		FromUser:   fromUser.ID,
		ToUser:     toUser.ID,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     model.StatusLent,
	}
	//when
	persistedLoanBook, err := repository.Persist(loanBook)
	//then
	assert.NoError(t, err)
	assert.NotNil(t, persistedLoanBook)
	assert.Equal(t, loanBook.LentAt, persistedLoanBook.LentAt)
	assert.NotZero(t, persistedLoanBook.ID)
}

func TestRepository_Persist_withInvalidFromUser(t *testing.T) {
	defer clearTable()
	//given
	loanBook := model.LoanBook{
		Book:       book.ID,
		FromUser:   99999,
		ToUser:     toUser.ID,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     model.StatusLent,
	}
	//when
	persistedLoanBook, err := repository.Persist(loanBook)
	//then
	assert.Nil(t, persistedLoanBook)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`book-lender`.`loan_books`, CONSTRAINT `fk_users_lent_books` FOREIGN KEY (`from_user`) REFERENCES `users` (`id`))")
}

func TestRepository_Persist_withInvalidBook(t *testing.T) {
	defer clearTable()
	//given
	loanBook := model.LoanBook{
		Book:       2345234,
		FromUser:   fromUser.ID,
		ToUser:     toUser.ID,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     model.StatusLent,
	}
	//when
	persistedLoanBook, err := repository.Persist(loanBook)
	//then
	assert.Nil(t, persistedLoanBook)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`book-lender`.`loan_books`, CONSTRAINT `fk_books_book_history` FOREIGN KEY (`book`) REFERENCES `books` (`id`))")
}

func TestRepository_Persist_withInvalidToUser(t *testing.T) {
	defer clearTable()
	//given
	loanBook := model.LoanBook{
		Book:       book.ID,
		FromUser:   fromUser.ID,
		ToUser:     1234,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     model.StatusLent,
	}
	//when
	persistedLoanBook, err := repository.Persist(loanBook)
	//then
	assert.Nil(t, persistedLoanBook)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`book-lender`.`loan_books`, CONSTRAINT `fk_users_borrowed_books` FOREIGN KEY (`to_user`) REFERENCES `users` (`id`))")
}

func TestRepository_Persist_withExistentId(t *testing.T) {
	defer clearTable()
	//given
	loanBook := model.LoanBook{
		Book:       book.ID,
		FromUser:   fromUser.ID,
		ToUser:     toUser.ID,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     model.StatusLent,
	}
	//when
	persistedLoanBook, err := repository.Persist(loanBook)
	assert.NoError(t, err)
	assert.NotNil(t, persistedLoanBook)

	persistedLoanBook2, err := repository.Persist(*persistedLoanBook)
	//then
	assert.Nil(t, persistedLoanBook2)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("Error 1062: Duplicate entry '%x' for key 'loan_books.PRIMARY'", persistedLoanBook.ID))
}

func TestRepository_Update(t *testing.T) {
	defer clearTable()
	//given
	loanBook := model.LoanBook{
		Book:       book.ID,
		FromUser:   fromUser.ID,
		ToUser:     toUser.ID,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     model.StatusLent,
	}

	persistedLoanBook, err := repository.Persist(loanBook)
	assert.NoError(t, err)
	assert.NotNil(t, persistedLoanBook)

	persistedLoanBook.Status = model.StatusReturned
	//when
	updated, err := repository.Update(*persistedLoanBook)
	//then
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, model.StatusReturned, updated.Status)
}

func TestRepository_Update_notFound(t *testing.T) {
	//given
	loanBook := model.LoanBook{
		ID:         1000,
		Book:       book.ID,
		FromUser:   fromUser.ID,
		ToUser:     toUser.ID,
		LentAt:     time.Now(),
		ReturnedAt: time.Now(),
		Status:     model.StatusLent,
	}
	//when
	updated, err := repository.Update(loanBook)
	//then
	assert.Error(t, err)
	assert.Nil(t, updated)
	assert.EqualError(t, err, "not found")
}
