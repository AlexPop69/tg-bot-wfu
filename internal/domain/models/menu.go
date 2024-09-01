package models

// MenuItem представляет собой позицию (товар) в меню.
type MenuItem struct {
	ID    int     // Уникальный идентификатор товара
	Name  string  // Название товара
	Price float64 // Цена товара
}
