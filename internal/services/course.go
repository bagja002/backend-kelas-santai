package services

import (
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/repository"

	"github.com/google/uuid"
)

type CourseService interface {
	CreateCourse(course *models.Course) error
	GetAllCourse(category string, status string) ([]models.Course, error)
	GetCourseByID(id uuid.UUID) (*models.Course, error)
	UpdateCourse(course *models.Course) error
	DeleteCourse(id uuid.UUID) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{
		repo: repo,
	}
}

func (s *courseService) CreateCourse(course *models.Course) error {
	return s.repo.CreateCourse(course)
}

func (s *courseService) GetAllCourse(category string, status string) ([]models.Course, error) {

	return s.repo.GetAllCourse(category, status)
}

func (s *courseService) GetCourseByID(id uuid.UUID) (*models.Course, error) {
	return s.repo.GetCourseByID(id)
}

func (s *courseService) UpdateCourse(course *models.Course) error {
	return s.repo.UpdateCourse(course)
}

func (s *courseService) DeleteCourse(id uuid.UUID) error {
	return s.repo.DeleteCourse(id)
}
