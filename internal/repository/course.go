package repository

import (
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course *models.Course) error
	GetAllCourse(category string, status string) ([]models.Course, error)
	GetCourseByID(id uuid.UUID) (*models.Course, error)
	UpdateCourse(course *models.Course) error
	DeleteCourse(id uuid.UUID) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{
		db: database.DB,
	}
}

func (r *courseRepository) CreateCourse(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) GetAllCourse(category string, status string) ([]models.Course, error) {
	var course []models.Course
	query := r.db.Model(&models.Course{})

	if category != "" {
		query = query.Where("category = ?", category)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	// http://localhost:4000/api/v1/static/public/uploads/courses/1767669526921117000.png
	err := query.Find(&course).Error
	baseUrl := "http://localhost:4000/api/v1/static/"

	for i, _ := range course {
		course[i].Picture = baseUrl + course[i].Picture
	}
	return course, err
}

func (r *courseRepository) GetCourseByID(id uuid.UUID) (*models.Course, error) {
	var course models.Course
	err := r.db.Preload("Curiculum").First(&course, "id = ?", id).Error
	baseUrl := "http://localhost:4000/api/v1/static/"
	course.Picture = baseUrl + course.Picture
	return &course, err
}

func (r *courseRepository) UpdateCourse(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) DeleteCourse(id uuid.UUID) error {
	return r.db.Delete(&models.Course{}, "id = ?", id).Error
}
