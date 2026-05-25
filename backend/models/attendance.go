package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	StudentID uuid.UUID      `gorm:"type:uuid;not null" json:"student_id"`
	Student   Student        `gorm:"foreignKey:StudentID" json:"student"`
	Date      time.Time      `gorm:"type:date;not null" json:"date"`
	Status    string         `gorm:"type:varchar(50);not null" json:"status"`
	Note      string         `gorm:"type:text" json:"note"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
