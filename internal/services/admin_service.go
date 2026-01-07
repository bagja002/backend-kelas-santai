package services

import (
	"errors"
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/repository"
	"project-kelas-santai/pkg/utils"

	"github.com/google/uuid"
)

type AdminService interface {
	CreateAdmin(admin *models.Admin) error
	Login(email, password string) (string, error)
	GetAdminByID(id uuid.UUID) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(id uuid.UUID) error
}

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{
		repo: repo,
	}
}

func (s *adminService) CreateAdmin(admin *models.Admin) error {
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return err
	}
	admin.Password = hashedPassword
	return s.repo.Create(admin)
}

func (s *adminService) Login(email, password string) (string, error) {
	admin, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(admin.ID.String(), "admin")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *adminService) GetAdminByID(id uuid.UUID) (*models.Admin, error) {
	return s.repo.FindByID(id)
}

func (s *adminService) UpdateAdmin(admin *models.Admin) error {
	return s.repo.Update(admin)
}

func (s *adminService) DeleteAdmin(id uuid.UUID) error {
	return s.repo.Delete(id)
}
