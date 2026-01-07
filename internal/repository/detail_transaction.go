package repository

import (
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type detailRepository struct {
	db *gorm.DB
}

type DetailTransactionRepository interface {
	Create(detail *models.DetailTransaction) error
	FindAll() ([]models.DetailTransaction, error)
	FindByID(id uuid.UUID) (*models.DetailTransaction, error)
	Update(detail *models.DetailTransaction) error
	Delete(id uuid.UUID) error
}

func NewDetailTransactionRepository() DetailTransactionRepository {
	return &detailRepository{
		db: database.DB,
	}
}

func (r *detailRepository) Create(detail *models.DetailTransaction) error {
	return r.db.Create(detail).Error
}

func (r *detailRepository) FindAll() ([]models.DetailTransaction, error) {
	var details []models.DetailTransaction
	err := r.db.Find(&details).Error
	return details, err
}

func (r *detailRepository) FindByID(id uuid.UUID) (*models.DetailTransaction, error) {
	var detail models.DetailTransaction
	err := r.db.First(&detail, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *detailRepository) Update(detail *models.DetailTransaction) error {
	return r.db.Save(detail).Error
}

func (r *detailRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.DetailTransaction{}, "id = ?", id).Error
}
