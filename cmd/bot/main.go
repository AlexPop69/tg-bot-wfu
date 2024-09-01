package main

import (
	"github.com/AlexPop69/tg-bot-wfu/internal/repository"
	"github.com/AlexPop69/tg-bot-wfu/internal/telegram"
	"github.com/AlexPop69/tg-bot-wfu/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	// Инициализация конфигурации, загрузка переменных окружения и настройка логирования
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Подключение к базе данных PostgreSQL
	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("failed to initialize database: %s", err.Error())
	}
	defer db.Close()

	// Инициализация репозиториев и сервисов

	// Инициализация Telegram-бота с передачей всех сервисов
	bot, err := telegram.NewBot(services)
	if err != nil {
		logrus.Fatalf("error occurred while initializing the bot: %s", err.Error())
	}

	// Запуск бота
	logrus.Println("Starting Telegram bot...")
	if err := bot.Start(); err != nil {
		logrus.Fatalf("error occurred while running the bot: %s", err.Error())
	}
}
