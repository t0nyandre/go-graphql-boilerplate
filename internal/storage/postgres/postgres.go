package postgres

import (
	"fmt"
	"os"
	"time"

	"github.com/t0nyandre/go-graphql-boilerplate/internal/user"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresConnect returns a new database instance
func NewPostgresConnect(logger *zap.SugaredLogger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Oslo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)), &gorm.Config{})

	if err != nil {
		logger.Warnf("Connection failed: %s. Retry connection in 5 seconds.", err.Error())
		time.Sleep(time.Duration(5) * time.Second)
		return NewPostgresConnect(logger)
	}

	db.AutoMigrate(&user.User{})

	logger.Info(fmt.Sprintf("Successfully connected to PostgreSQL with database: %s", os.Getenv("POSTGRES_DATABASE")))

	return db
}
