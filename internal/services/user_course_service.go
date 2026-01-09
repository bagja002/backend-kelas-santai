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
	GetUserCourseDashboard(userID uuid.UUID) ([]models.UserCourseDashboardResponse, error)
	GetCoursePending(userID uuid.UUID, status string) ([]models.PendingCourse, error)
	UpdateProgress(userID, courseID uuid.UUID, progress int) error
	Delete(userID, courseID uuid.UUID) error
}

type TransactionService interface {
	PaymentCourse(userId uuid.UUID, courseId []uuid.UUID) (error, string)
	GetTransactionHistory(userID uuid.UUID) ([]models.TransactionHistoryResponse, error)
	HandleNotification(payload map[string]interface{}) error
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
	return s.repo.FindByUser(userID, "paid")
}

func (s *userCourseService) GetUserCourseDashboard(userID uuid.UUID) ([]models.UserCourseDashboardResponse, error) {
	userCourses, err := s.repo.FindDashboardByUser(userID)
	if err != nil {
		return nil, err
	}

	var response []models.UserCourseDashboardResponse
	for _, uc := range userCourses {
		item := models.UserCourseDashboardResponse{
			ID:       uc.ID,
			CourseID: uc.CourseID,
			Status:   uc.Status,
			Progress: uc.Progress,
			Course: models.CourseDashboardItem{
				ID:           uc.Course.ID,
				Title:        uc.Course.Title,
				Image:        s.cfg.Web.BaseUrl + uc.Course.Picture,
				Picture:      s.cfg.Web.BaseUrl + uc.Course.Picture,
				Level:        uc.Course.Level,
				MentorName:   uc.Course.MentorName,
				TotalModules: len(uc.Course.Curiculum),
			},
		}
		response = append(response, item)
	}

	return response, nil
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

	return nil, snapResp.RedirectURL
}

func (s *transactionService) GetTransactionHistory(userID uuid.UUID) ([]models.TransactionHistoryResponse, error) {
	transactions, err := s.repo.FindByUser(userID)
	if err != nil {
		return nil, err
	}

	var history []models.TransactionHistoryResponse
	for _, t := range transactions {
		var items []models.TransactionItem
		totalAmount := 0.0

		for _, detail := range t.DetailTransaction {
			items = append(items, models.TransactionItem{
				CourseTitle: detail.Course.Title,
				Price:       detail.Course.Price,
			})
			totalAmount += detail.Course.Price
		}

		history = append(history, models.TransactionHistoryResponse{
			ID:        t.ID.String(),
			Code:      t.PaymentID,
			Amount:    totalAmount,
			Status:    t.Status,
			CreatedAt: t.CreatedAt,
			Items:     items,
		})
	}

	return history, nil
}

func (s *transactionService) HandleNotification(payload map[string]interface{}) error {
	orderID, ok := payload["order_id"].(string)
	if !ok {
		return nil // or error
	}

	//order ID ==Payent Id
	fmt.Println("orderID", orderID)

	transaction, err := s.repo.FindByPaymentID(orderID)
	if err != nil {
		return err
	}

	transactionStatus, ok := payload["transaction_status"].(string)
	if !ok {
		return nil // or error
	}

	fraudStatus, _ := payload["fraud_status"].(string)

	var newStatus string

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			newStatus = "challenge"
		} else if fraudStatus == "accept" {
			newStatus = "paid"
		}
	} else if transactionStatus == "settlement" {
		newStatus = "paid"
	} else if transactionStatus == "deny" {
		newStatus = "cancelled"
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		newStatus = "cancelled"
	} else if transactionStatus == "pending" {
		newStatus = "pending"
	}
	if newStatus != "" {
		transaction.Status = newStatus
		s.repo.Update(transaction)
	}
	if transaction.Status == "paid" {
		//Cari Detail Transaction
		detailTransaction, err := s.repo.FindAllDetailTransaction(transaction.ID)
		if err != nil {
			return err
		}

		iduser := transaction.UserID

		for _, detail := range detailTransaction {
			courseId := detail.CourseID
			s.repo.UpdatePaidUserCourse(iduser, courseId)
		}

		fmt.Println("detailTransaction", detailTransaction)

	}

	return nil
}
