package service

import (
	"github.com/AlexPop69/tg-bot-wfu/internal/domain/models"
	"github.com/AlexPop69/tg-bot-wfu/internal/repository"
)

// AdminService предоставляет методы для управления администраторами.
type AdminService struct {
	repository *repository.Repository // Слой доступа к данным
}

// NewAdminService создает новый экземпляр AdminService.
func NewAdminService(repository *repository.Repository) *AdminService {
	return &AdminService{repository: repository}
}

// Authentication выполняет аутентификацию администратора.
func (s *AdminService) Authentication(username string) (*models.Admin, error) {
	return s.repository.GetAdminByUsername(username)
}
