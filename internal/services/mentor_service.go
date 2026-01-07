package services

import (
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/repository"

	"github.com/google/uuid"
)

type MentorService interface {
	CreateMentor(mentor *models.Mentor) error
	GetAllMentors() ([]models.Mentor, error)
	GetMentorByID(id uuid.UUID) (*models.Mentor, error)
	UpdateMentor(mentor *models.Mentor) error
	DeleteMentor(id uuid.UUID) error
}

type mentorService struct {
	repo repository.MentorRepository
}

func NewMentorService(repo repository.MentorRepository) MentorService {
	return &mentorService{
		repo: repo,
	}
}

func (s *mentorService) CreateMentor(mentor *models.Mentor) error {
	return s.repo.Create(mentor)
}

func (s *mentorService) GetAllMentors() ([]models.Mentor, error) {
	return s.repo.FindAll()
}

func (s *mentorService) GetMentorByID(id uuid.UUID) (*models.Mentor, error) {
	return s.repo.FindByID(id)
}

func (s *mentorService) UpdateMentor(mentor *models.Mentor) error {
	return s.repo.Update(mentor)
}

func (s *mentorService) DeleteMentor(id uuid.UUID) error {
	return s.repo.Delete(id)
}
