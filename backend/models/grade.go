package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Grade struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	StudentID    uuid.UUID      `gorm:"type:uuid;not null" json:"student_id"`
	Student      Student        `gorm:"foreignKey:StudentID" json:"student"`
	Subject      string         `gorm:"type:varchar(255);not null" json:"subject"`
	Score        float64        `gorm:"type:decimal(5,2);not null" json:"score"`
	Semester     int            `gorm:"type:integer;not null" json:"semester"`
	AcademicYear string         `gorm:"type:varchar(50);not null" json:"academic_year"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
