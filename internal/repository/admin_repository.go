package repository

import (
	"fmt"

	"github.com/AlexPop69/tg-bot-wfu/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

// GetAdminByUsername метод для получения админа по его telegram username
func (s *AdminPostgres) GetAdminByUsername(username string) (*models.Admin, error) {
	admin := &models.Admin{}

	query := fmt.Sprintf(`SELECT id, telegramUsername
		FROM %s
		WHERE telegramUsername = $1`, adminsTable)

	// Выполняем запрос на выборку данных об администраторе по его username.
	err := s.db.QueryRow(query, username).Scan(&admin.ID, &admin.TelegramUsername)
	if err != nil {
		return nil, fmt.Errorf("GetAdminByUsername: Admin %s does not exist", username)
	}

	return admin, nil
}
