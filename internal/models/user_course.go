package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCourse struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	CourseID  uuid.UUID `gorm:"type:char(36)" json:"course_id"`
	Status    string    `json:"status"`
	Progress  int       `json:"progress"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type PendingCourse struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Mentor string    `json:"mentor"`
	Price  float64   `json:"price"`
	Image  string    `json:"image"`
}

func (uc *UserCourse) BeforeCreate(tx *gorm.DB) (err error) {
	uc.ID = uuid.New()
	return
}
