package repository

import (
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindAll() ([]models.Transaction, error)
	FindByID(id uuid.UUID) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id uuid.UUID) error
	FindUser(id uuid.UUID) (*models.User, error)
	FindCourse(id uuid.UUID) (*models.Course, error)
	CreateDetailTransaction(detail *models.DetailTransaction) error
	FindAllDetailTransaction() ([]models.DetailTransaction, error)
	FindByIDDetailTransaction(id uuid.UUID) (*models.DetailTransaction, error)
	UpdateDetailTransaction(detail *models.DetailTransaction) error
	DeleteDetailTransaction(id uuid.UUID) error
}

type transactionRepository struct {
	db *gorm.DB
}

func (r *transactionRepository) FindUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{
		db: database.DB,
	}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindCourse(id uuid.UUID) (*models.Course, error) {
	var course models.Course
	err := r.db.First(&course, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *transactionRepository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) FindByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Transaction{}, "id = ?", id).Error
}

func (r *transactionRepository) CreateDetailTransaction(detail *models.DetailTransaction) error {
	return r.db.Create(detail).Error
}

func (r *transactionRepository) FindAllDetailTransaction() ([]models.DetailTransaction, error) {
	var details []models.DetailTransaction
	err := r.db.Find(&details).Error
	return details, err
}

func (r *transactionRepository) FindByIDDetailTransaction(id uuid.UUID) (*models.DetailTransaction, error) {
	var detail models.DetailTransaction
	err := r.db.First(&detail, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *transactionRepository) UpdateDetailTransaction(detail *models.DetailTransaction) error {
	return r.db.Save(detail).Error
}

func (r *transactionRepository) DeleteDetailTransaction(id uuid.UUID) error {
	return r.db.Delete(&models.DetailTransaction{}, "id = ?", id).Error
}
