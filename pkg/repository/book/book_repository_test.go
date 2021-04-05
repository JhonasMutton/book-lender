package book

import (
	"github.com/JhonasMutton/book-lender/pkg/database"
	"github.com/JhonasMutton/book-lender/pkg/model"
	testDatabase "github.com/JhonasMutton/book-lender/pkg/test/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

var genericUser model.User
var db *gorm.DB
var repository *Repository

func init() {
	testDatabase.InitMySqlContainer()
	setupTest()
}

func setupTest() {
	db = database.NewDatabaseConnection()
	repository = NewRepository(db)

	genericUser = model.User{
		Name:      "Bruce Wayne",
		Email:     "bruce.wayne@wayne.com",
		CreatedAt: time.Now(),
	}

	result := db.Create(&genericUser)
	if err := result.Error; err != nil {
		panic("error to create user" + err.Error())
	}
}

func TestNewRepository(t *testing.T) {
	repository := NewRepository(db)

	assert.NotNil(t, repository)
	assert.NotNil(t, repository.db)
	assert.Equal(t, db, repository.db)
}

func TestRepository_Persist(t *testing.T) {
	//given
	book := model.Book{
		Title:     "A volta dos que não foram",
		Pages:     760,
		CreatedAt: time.Now(),
		OwnerId:   genericUser.ID,
	}
	//when
	persistedBook, err := repository.Persist(book)
	//then
	assert.NoError(t, err)
	assert.NotNil(t, persistedBook)
	assert.Equal(t, book.CreatedAt, persistedBook.CreatedAt)
	assert.NotZero(t, persistedBook.ID)
}

func TestRepository_Persist_withoutOwner(t *testing.T) {
	//given
	book := model.Book{
		Title:     "A volta dos que não foram",
		Pages:     760,
		CreatedAt: time.Now(),
	}
	//when
	persistedBook, err := repository.Persist(book)
	//then
	assert.Nil(t, persistedBook)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`book-lender`.`books`, CONSTRAINT `fk_users_collection` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`))")
}

func TestRepository_Persist_withInvalidOwner(t *testing.T) {
	//given
	book := model.Book{
		Title:     "A volta dos que não foram",
		Pages:     760,
		CreatedAt: time.Now(),
		OwnerId: 10,
	}
	//when
	persistedBook, err := repository.Persist(book)
	//then
	assert.Nil(t, persistedBook)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`book-lender`.`books`, CONSTRAINT `fk_users_collection` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`))")
	//TODO adicionar validação comparando com classes de error da lib
}

func TestRepository_Persist_withExistentId(t *testing.T) {
	//given
	book := model.Book{
		Title:     "A volta dos que não foram",
		Pages:     760,
		CreatedAt: time.Now(),
		OwnerId: genericUser.ID,
	}
	//when
	persistedBook, err := repository.Persist(book)
	assert.NoError(t, err)
	assert.NotNil(t, persistedBook)
	assert.NotZero(t, persistedBook.ID)

	persistedBook2, err := repository.Persist(*persistedBook)

	//then
	assert.Nil(t, persistedBook2)
	assert.Error(t, err)
	assert.EqualError(t, err, "Error 1062: Duplicate entry '1' for key 'books.PRIMARY'")
}