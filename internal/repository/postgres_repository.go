package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL драйвер
	"github.com/spf13/viper"
)

const (
	adminsTable = "admins"
	shopsTable  = "shops"
	usersTable  = "users"
	menuTable   = "menu"
	ordersTable = "orders"
)

// NewPostgresDB создает новое подключение к базе данных PostgreSQL
func NewPostgresDB() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		os.Getenv("POSTGRES_PASSWORD"),
		viper.GetString("database.dbname"),
		viper.GetString("database.sslmode"),
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the PostgreSQL database successfully")
	return db, nil
}
