package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"type:varchar(191);uniqueIndex" json:"email"`
	NoTelp   string    `json:"no_telp"`
	Address  string    `json:"address"`
	City     string    `json:"city"`
	Province string    `json:"province"`
	Gender   string    `json:"gender"`

	Password  string    `json:"-"` // Don't return password in JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
