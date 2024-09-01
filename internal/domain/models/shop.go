package models

// Shop представляет собой магазин.
type Shop struct {
	ID        int    // Уникальный идентификатор магазина
	Name      string // Название магазина
	Address   string // Адрес магазина
	OpenTime  string // Время открытия магазина (в формате HH:MM)
	CloseTime string // Время закрытия магазина (в формате HH:MM)
}
