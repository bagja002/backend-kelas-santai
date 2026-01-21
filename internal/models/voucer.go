package models

import "github.com/google/uuid"

type Voucer struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `json:"name"`
	Discount  uint
	CreateAt  string
	UpdatedAt string
}
