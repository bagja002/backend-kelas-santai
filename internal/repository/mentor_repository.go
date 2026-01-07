package repository

import (
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MentorRepository interface {
	Create(mentor *models.Mentor) error
	FindAll() ([]models.Mentor, error)
	FindByID(id uuid.UUID) (*models.Mentor, error)
	Update(mentor *models.Mentor) error
	Delete(id uuid.UUID) error
}

type mentorRepository struct {
	db *gorm.DB
}

func NewMentorRepository() MentorRepository {
	return &mentorRepository{
		db: database.DB,
	}
}

func (r *mentorRepository) Create(mentor *models.Mentor) error {
	return r.db.Create(mentor).Error
}

func (r *mentorRepository) FindAll() ([]models.Mentor, error) {
	var mentors []models.Mentor
	err := r.db.Find(&mentors).Error
	return mentors, err
}

func (r *mentorRepository) FindByID(id uuid.UUID) (*models.Mentor, error) {
	var mentor models.Mentor
	err := r.db.First(&mentor, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &mentor, nil
}

func (r *mentorRepository) Update(mentor *models.Mentor) error {
	return r.db.Save(mentor).Error
}

func (r *mentorRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Mentor{}, "id = ?", id).Error
}
