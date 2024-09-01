package models

import "time"

// Order представляет собой заказ.
type Order struct {
	ID          int            // Уникальный идентификатор заказа
	User        string         // Telegram username пользователя, сделавшего заказ
	Shop        string         // Название магазина, где был сделан заказ
	Items       map[string]int // Список товаров и их количества в заказе (используется для JSONB поля)
	CreatedAt   time.Time      // Дата и время создания заказа
	OrderAmount float64        // Общая сумма заказа
}
