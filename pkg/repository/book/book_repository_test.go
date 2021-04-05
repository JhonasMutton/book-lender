package book

import (
	"github.com/JhonasMutton/book-lender/pkg/test/database"
	"testing"
)

func init() {
	database.InitMySqlContainer()
}

func TestName(t *testing.T) {
	print("show")
}
