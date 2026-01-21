package services

import (
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/repository"

	"github.com/google/uuid"
)

type VoucerService interface {
	CreateVoucer(voucer *models.Voucer) error
	GetVoucerAll() ([]models.Voucer, error)
	GetVoucerById(id uuid.UUID) (*models.Voucer, error)
	UpdateVoucer(voucer *models.Voucer) error
	DeleteVoucer(id string) error
	GetVoucerName(name string) (models.Voucer, error)
}

type voucerService struct {
	repo repository.VoucerRepository
}

func NewVoucerService(repo repository.VoucerRepository) VoucerService {
	return &voucerService{
		repo: repo,
	}
}

func (s *voucerService) CreateVoucer(voucer *models.Voucer) error {
	return s.repo.CreateVoucer(voucer)
}

func (s *voucerService) GetVoucerAll() ([]models.Voucer, error) {
	return s.repo.GetVoucerAll()
}

func (s *voucerService) GetVoucerById(id uuid.UUID) (*models.Voucer, error) {
	return s.repo.GetVoucerById(id)
}

func (s *voucerService) UpdateVoucer(voucer *models.Voucer) error {
	return s.repo.UpdateVoucer(voucer)
}

func (s *voucerService) DeleteVoucer(id string) error {
	return s.repo.DeleteVoucer(id)
}

func (s *voucerService) GetVoucerName(name string) (models.Voucer, error) {
	return s.repo.GetVoucerName(name)
}
