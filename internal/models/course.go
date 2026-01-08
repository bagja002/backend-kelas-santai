package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID          uuid.UUID   `gorm:"type:char(36);primaryKey" json:"id"`
	MentorID    string      `json:"mentor_id"`
	Name        string      `json:"name"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Level       string      `json:"level"`
	Rating      float64     `json:"rating"`
	Category    string      `json:"category"`
	Picture     string      `json:"picture"`
	Duration    string      `json:"duration"`
	TotalJp     int         `json:"total_jp"`
	MentorName  string      `json:"mentor_name"`
	StartDate   string      `json:"start_date"`
	EndDate     string      `json:"end_date"`
	Silabus     string      `json:"silabus"`
	GarisBesar  string      `json:"garis_besar"`
	Status      string      `json:"status"`
	Curiculum   []Curiculum `json:"curiculum" gorm:"foreignKey:CourseID;references:ID"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

type Curiculum struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	CourseID    uuid.UUID `gorm:"type:char(36)" json:"course_id"`
	NoUrut      int       `json:"no_urut"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

func (c *Curiculum) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
