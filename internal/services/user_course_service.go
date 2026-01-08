package services

import (
	"errors"
	"fmt"
	"project-kelas-santai/internal/config"
	"project-kelas-santai/internal/models"
	"project-kelas-santai/internal/repository"
	"project-kelas-santai/pkg/utils"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type UserCourseService interface {
	EnrollCourse(userID, courseID uuid.UUID) error
	GetUserCourses(userID uuid.UUID) ([]models.UserCourse, error)
	GetCoursePending(userID uuid.UUID, status string) ([]models.PendingCourse, error)
	UpdateProgress(userID, courseID uuid.UUID, progress int) error
	Delete(userID, courseID uuid.UUID) error
}

type TransactionService interface {
	PaymentCourse(userId uuid.UUID, courseId []uuid.UUID) (error, string)
}

type userCourseService struct {
	repo repository.UserCourseRepository
	cfg  *config.Config
}

type transactionService struct {
	repo       repository.TransactionRepository
	snapClient snap.Client
	config     *config.Config
}

func NewUserCourseService(repo repository.UserCourseRepository, cfg *config.Config) UserCourseService {
	return &userCourseService{
		repo: repo,
		cfg:  cfg,
	}
}

func NewTransactionService(repo repository.TransactionRepository, config *config.Config) TransactionService {

	var env midtrans.EnvironmentType
	if config.Midtrans.Environment == "production" {
		env = midtrans.Production
	} else {
		env = midtrans.Sandbox
	}

	var s snap.Client
	s.New(config.Midtrans.ServerKey, env)
	return &transactionService{
		repo:       repo,
		config:     config,
		snapClient: s,
	}
}

func (s *userCourseService) EnrollCourse(userID, courseID uuid.UUID) error {
	existing, _ := s.repo.FindOne(userID, courseID)
	if existing != nil {
		return errors.New("user already enrolled in this course")
	}

	userCourse := &models.UserCourse{
		UserID:   userID,
		CourseID: courseID,
		Status:   "pending", //karna status blom di bayar
		Progress: 0,
	}

	return s.repo.Create(userCourse)
}

func (s *userCourseService) GetUserCourses(userID uuid.UUID) ([]models.UserCourse, error) {
	return s.repo.FindByUser(userID, "aktif")
}

func (s *userCourseService) GetCoursePending(userID uuid.UUID, status string) ([]models.PendingCourse, error) {
	courses, err := s.repo.GetPendingCourses(userID)
	if err != nil {
		return nil, err
	}

	// Format Image URL
	baseUrl := s.cfg.Web.BaseUrl
	for i := range courses {
		courses[i].Image = baseUrl + courses[i].Image
	}

	return courses, nil
}

func (s *userCourseService) UpdateProgress(userID, courseID uuid.UUID, progress int) error {
	userCourse, err := s.repo.FindOne(userID, courseID)
	if err != nil {
		return err
	}

	userCourse.Progress = progress
	if progress >= 100 {
		userCourse.Status = "completed"
	}

	return s.repo.Update(userCourse)
}

func (s *userCourseService) Delete(userID, courseID uuid.UUID) error {
	userCourse, err := s.repo.FindOne(userID, courseID)
	if err != nil {
		return err
	}
	return s.repo.Delete(userCourse)
}

func (s *transactionService) PaymentCourse(userId uuid.UUID, courseId []uuid.UUID) (error, string) {

	//Create Transcation Data
	transaction := &models.Transaction{
		UserID:    userId,
		Status:    "pending",
		PaymentID: utils.GeneratorTransactionId(),
		CreatedAt: utils.TimeNowJakarta(),
		UpdatedAt: utils.TimeNowJakarta(),
	}

	if err := s.repo.Create(transaction); err != nil {
		return err, ""
	}
	user, err := s.repo.FindUser(userId)
	if err != nil {
		return err, ""
	}

	//MEndapatka Total Harga Course
	totalPrice := 0
	for _, courseID := range courseId {
		course, err := s.repo.FindCourse(courseID)
		if err != nil {
			return err, ""
		}
		totalPrice += int(course.Price)
	}

	//Create Detail Transcation
	for _, courseID := range courseId {
		transactionDetail := &models.DetailTransaction{
			TransactionID: transaction.ID,
			CourseID:      courseID,
			CreatedAt:     utils.TimeNowJakarta(),
			UpdatedAt:     utils.TimeNowJakarta(),
		}
		if err := s.repo.CreateDetailTransaction(transactionDetail); err != nil {
			return err, ""
		}
	}

	//Integrasi Midtrans
	snapReq := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.PaymentID,
			GrossAmt: int64(totalPrice),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
		},
	}

	snapResp, snapErr := s.snapClient.CreateTransaction(&snapReq)
	if snapErr != nil {
		return snapErr, ""
	}
	fmt.Println(snapResp)

	return nil, snapResp.RedirectURL
}
