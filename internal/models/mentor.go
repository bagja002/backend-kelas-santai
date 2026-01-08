package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mentor struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name          string    `json:"name"`
	Avatar        string    `json:"avatar"`
	Email         string    `gorm:"type:varchar(191);uniqueIndex"`
	Password      string    `json:"-"` // Don't return password in JSON
	GelarDepan    string    `json:"gelar_depan"`
	GelarBelakang string    `json:"gelar_belakang"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

func (m *Mentor) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
