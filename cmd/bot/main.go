package main

import (
	"log"

	"well-fed_uncle/internal/repository"
	"well-fed_uncle/pkg/config"
)

func main() {
	// Инициализация конфигурации, загрузка переменных окружения и настройка логирования
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	// Подключение к базе данных PostgreSQL
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}
	defer db.Close()

	// Инициализация репозиториев и сервисов

	// Инициализация Telegram-бота с передачей всех сервисов

	// Запуск бота

}
