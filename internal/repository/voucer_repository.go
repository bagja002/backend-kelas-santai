package repository

import (
	"project-kelas-santai/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type voucerRepository struct {
	db *gorm.DB
}

func (v *voucerRepository) CreateVoucer(voucer *models.Voucer) error {
	return v.db.Create(voucer).Error
}

func (v *voucerRepository) GetVoucerAll() ([]models.Voucer, error) {
	var voucers []models.Voucer
	err := v.db.Find(&voucers).Error
	return voucers, err
}

func (v *voucerRepository) GetVoucerById(id uuid.UUID) (*models.Voucer, error) {
	var voucer models.Voucer
	err := v.db.Where("id = ?", id).First(&voucer).Error
	return &voucer, err
}

func (v *voucerRepository) UpdateVoucer(voucer *models.Voucer) error {
	return v.db.Save(voucer).Error
}

func (v *voucerRepository) DeleteVoucer(id string) error {
	return v.db.Where("id = ?", id).Delete(&models.Voucer{}).Error
}

type VoucerRepository interface {
	CreateVoucer(voucer *models.Voucer) error
	GetVoucerAll() ([]models.Voucer, error)
	GetVoucerById(id uuid.UUID) (*models.Voucer, error)
	UpdateVoucer(voucer *models.Voucer) error
	DeleteVoucer(id string) error
}

func NewVoucerRepository(db *gorm.DB) VoucerRepository {
	return &voucerRepository{db: db}
}
