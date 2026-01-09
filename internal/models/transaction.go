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
	Course        Course    `gorm:"foreignKey:CourseID" json:"course"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

type TransactionHistoryResponse struct {
	ID        string            `json:"id"`
	Code      string            `json:"code"`
	Amount    float64           `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt string            `json:"created_at"`
	Items     []TransactionItem `json:"items"`
}

type TransactionItem struct {
	CourseTitle string  `json:"course_title"`
	Price       float64 `json:"price"`
}

func (m *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

func (m *DetailTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
