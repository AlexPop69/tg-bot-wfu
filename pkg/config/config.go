package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitConfig загружает конфигурацию, переменные окружения и настраивает логирование
func InitConfig() error {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found")
	}

	// Настройка логирования
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	// Загрузка конфигурации из файла config.yml
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
