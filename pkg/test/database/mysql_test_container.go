package database

import (
	"context"
	"fmt"
	"github.com/JhonasMutton/book-lender/pkg/database"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitMySqlContainer() {
	ctx := context.Background()
	user := "root"
	password := "admin"
	dbName := "book-lender"

	req := testcontainers.ContainerRequest{
		Image:        "mysql:" + database.MySqlVersion,
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "admin",
			"MYSQL_DATABASE":      dbName,
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	mysqlC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")

	dsn := fmt.Sprintf(database.MySqlDsnFormat,
		user,
		password,
		host,
		p.Port(),
		dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}

	migration(db)

	SetupEnvVars(dbName, user, password, host, p.Port())
}

func migration(db *gorm.DB) {
	if err := db.AutoMigrate(model.User{}, model.LoanBook{}, model.Book{}); err != nil {
		panic("error to migration:" + err.Error())
	}
}

func SetupEnvVars(dbName, user, password, host, port string) {
	_ = os.Setenv("DB_NAME", dbName)
	_ = os.Setenv("DB_USER", user)
	_ = os.Setenv("DB_PASSWORD", password)
	_ = os.Setenv("DB_HOST", host)
	_ = os.Setenv("DB_PORT", port)
}
