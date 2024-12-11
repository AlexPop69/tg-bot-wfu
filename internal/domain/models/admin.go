package models

// Admin представляет собой администратора.
type Admin struct {
	ID               int    // Уникальный идентификатор администратора
	Name             string // Имя администратора
	TelegramUsername string // Telegram username администратора
	Shop             string // Название магазина, которым управляет администратор
	IsSuperAdmin     bool   // Является ли администратор суперадмином
}
