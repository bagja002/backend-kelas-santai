package repository

import (
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCourseRepository interface {
	Create(userCourse *models.UserCourse) error
	FindByUser(userID uuid.UUID, status string) ([]models.UserCourse, error)
	FindByCourse(courseID uuid.UUID) ([]models.UserCourse, error)
	FindOne(userID uuid.UUID, courseID uuid.UUID) (*models.UserCourse, error)
	GetPendingCourses(userID uuid.UUID) ([]models.PendingCourse, error)
	Update(userCourse *models.UserCourse) error
	Delete(userCourse *models.UserCourse) error
}

type userCourseRepository struct {
	db *gorm.DB
}

func NewUserCourseRepository() UserCourseRepository {
	return &userCourseRepository{
		db: database.DB,
	}
}

func (r *userCourseRepository) Create(userCourse *models.UserCourse) error {
	return r.db.Create(userCourse).Error
}

func (r *userCourseRepository) FindByUser(userID uuid.UUID, status string) ([]models.UserCourse, error) {
	var userCourses []models.UserCourse
	err := r.db.Where("user_id = ? AND status = ?", userID, status).Find(&userCourses).Error
	return userCourses, err
}

func (r *userCourseRepository) FindByCourse(courseID uuid.UUID) ([]models.UserCourse, error) {
	var userCourses []models.UserCourse
	err := r.db.Where("course_id = ?", courseID).Find(&userCourses).Error
	return userCourses, err
}

func (r *userCourseRepository) FindOne(userID uuid.UUID, courseID uuid.UUID) (*models.UserCourse, error) {
	var userCourse models.UserCourse
	err := r.db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&userCourse).Error
	if err != nil {
		return nil, err
	}
	return &userCourse, nil
}

func (r *userCourseRepository) Update(userCourse *models.UserCourse) error {
	return r.db.Save(userCourse).Error
}

func (r *userCourseRepository) Delete(userCourse *models.UserCourse) error {
	return r.db.Delete(userCourse).Error
}

func (r *userCourseRepository) GetPendingCourses(userID uuid.UUID) ([]models.PendingCourse, error) {
	var pendingCourses []models.PendingCourse
	err := r.db.Table("user_courses").
		Select("courses.id, courses.title, courses.mentor_name as mentor, courses.price, courses.picture as image").
		Joins("JOIN courses ON courses.id = user_courses.course_id").
		Where("user_courses.user_id = ? AND user_courses.status = ?", userID, "pending").
		Scan(&pendingCourses).Error
	return pendingCourses, err
}
