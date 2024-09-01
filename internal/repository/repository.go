package repository

import (
	"github.com/AlexPop69/tg-bot-wfu/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type Admin interface {
	GetAdminByUsername(username string) (*models.Admin, error)
}

type Repository struct {
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Admin: NewAdminPostgres(db),
	}
}
