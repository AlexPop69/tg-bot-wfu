package models

import "time"

// User представляет собой пользователя бота.
type User struct {
	ID               int       // Уникальный идентификатор пользователя
	Name             string    // Имя пользователя
	TelegramUsername string    // Telegram username пользователя
	PhoneNumber      string    // Номер телефона пользователя
	NumberOfOrders   int       // Количество заказов, сделанных пользователем
	AddedAt          time.Time // Дата и время первого добавления пользователя
}
