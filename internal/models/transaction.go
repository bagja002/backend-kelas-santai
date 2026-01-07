package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                uuid.UUID           `gorm:"type:char(36);primaryKey" json:"id"`
	UserID            uuid.UUID           `gorm:"type:char(36)" json:"user_id"`
	Status            string              `json:"status"`
	PaymentID         string              `json:"payment_id"`
	CreatedAt         string              `json:"created_at"`
	UpdatedAt         string              `json:"updated_at"`
	DetailTransaction []DetailTransaction `gorm:"foreignKey:TransactionID"`
}

type DetailTransaction struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	TransactionID uuid.UUID `gorm:"type:char(36)" json:"transaction_id"`
	CourseID      uuid.UUID `gorm:"type:char(36)" json:"course_id"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

func (m *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

func (m *DetailTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
