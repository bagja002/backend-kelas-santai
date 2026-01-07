package services

import (
	"errors"
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/repository"
	"project-kelas-santai/pkg/utils"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *models.User) error
	Login(email, password string) (string, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(id uuid.UUID, user *models.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (string, error) {

	if email == "" || password == "" {
		return "", errors.New("email or password is empty")
	}
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("Email not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("Password not match")
	}

	token, err := utils.GenerateToken(user.ID.String(), "user")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uuid.UUID, data *models.User) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	user.Name = data.Name
	user.Email = data.Email

	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}
