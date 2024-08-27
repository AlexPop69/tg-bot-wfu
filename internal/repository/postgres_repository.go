package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL драйвер
	"github.com/spf13/viper"
)

// NewPostgresDB создает новое подключение к базе данных PostgreSQL
func NewPostgresDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.dbname"),
		viper.GetString("database.sslmode"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the PostgreSQL database successfully")
	return db, nil
}
