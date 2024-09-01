package service

import (
	"github.com/AlexPop69/tg-bot-wfu/internal/domain/models"
	"github.com/AlexPop69/tg-bot-wfu/internal/repository"
)

type Admin interface {
	Authentication(username string) (*models.Admin, error)
}

// Service объединяет все сервисы
type Service struct {
	Admin
}

// NewService создает новый экземпляр Service
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Admin: NewAdminService(repo),
	}
}
