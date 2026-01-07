package repository

import (
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *models.Admin) error
	FindByEmail(email string) (*models.Admin, error)
	FindByID(id uuid.UUID) (*models.Admin, error)
	Update(admin *models.Admin) error
	Delete(id uuid.UUID) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{
		db: database.DB,
	}
}

func (r *adminRepository) Create(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) FindByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) FindByID(id uuid.UUID) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.First(&admin, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Update(admin *models.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Admin{}, "id = ?", id).Error
}
