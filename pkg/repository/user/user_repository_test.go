package user

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

var db *gorm.DB
var repository *Repository

func init() {
	testDatabase.InitMySqlContainer()
	setupTest()
}

func setupTest() {
	db = database.NewDatabaseConnection()
	repository = NewRepository(db)
}

func clearTable() {
	db.Exec("DELETE FROM users")
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
	user := model.User{
		Name:      "Bruce Wayne",
		Email:     "bruce.wayne@wayne.com",
		CreatedAt: time.Now(),
	}
	//when
	persistedUser, err := repository.Persist(user)
	//then
	assert.NoError(t, err)
	assert.NotNil(t, persistedUser)
	assert.Equal(t, user.CreatedAt, persistedUser.CreatedAt)
	assert.NotZero(t, persistedUser.ID)


}

func TestRepository_Persist_withExistentId(t *testing.T) {
	defer clearTable()
	//given
	user := model.User{
		Name:      "Bruce Wayne",
		Email:     "bruce.wayne@wayne.com",
		CreatedAt: time.Now(),
	}
	//when
	persistedUser, err := repository.Persist(user)
	assert.NoError(t, err)
	assert.NotNil(t, persistedUser)
	assert.NotZero(t, persistedUser.ID)

	persistedUser2, err := repository.Persist(*persistedUser)
	//then
	assert.Nil(t, persistedUser2)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("Error 1062: Duplicate entry '%x' for key 'users.PRIMARY'", persistedUser.ID))
}

func TestRepository_FetchById(t *testing.T) {
	defer clearTable()
	//given
	user := model.User{
		Name:      "Bruce Wayne",
		Email:     "bruce.wayne@wayne.com",
		CreatedAt: time.Now(),
	}

	persistedUser, err := repository.Persist(user)
	assert.NoError(t, err)
	assert.NotNil(t, persistedUser)
	//when
	userFound, err := repository.FetchById(persistedUser.ID)
	//then
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
	assert.NotZero(t, userFound.ID)
	assert.Equal(t, persistedUser.ID, userFound.ID)
}

func TestRepository_FetchByIdNotFound(t *testing.T) {
	defer clearTable()
	//given
	userId := uint(1939)
	//when
	userFound, err := repository.FetchById(userId)
	//then
	assert.Error(t, err)
	assert.EqualError(t, err, "record not found")
	assert.Nil(t, userFound)
}

func TestRepository_Fetch(t *testing.T) {
	defer clearTable()
	//given
	user := model.User{
		Name:      "Bruce Wayne",
		Email:     "bruce.wayne@wayne.com",
		CreatedAt: time.Now(),
	}

	persistedUser, err := repository.Persist(user)
	assert.NoError(t, err)
	assert.NotNil(t, persistedUser)
	//when
	usersFound, err := repository.Fetch()
	//then
	assert.NoError(t, err)
	assert.NotNil(t, usersFound)
	assert.NotEmpty(t, usersFound)
}

func TestRepository_FetchWithoutUsers(t *testing.T) {
	defer clearTable()
	//when
	usersFound, err := repository.Fetch()
	//then
	assert.NoError(t, err)
	assert.NotNil(t, usersFound)
	assert.Empty(t, usersFound)
}