package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	NISN      string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"nisn"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Class     string         `gorm:"type:varchar(50);not null" json:"class"`
	Gender    string         `gorm:"type:varchar(20);not null" json:"gender"`
	Address   string         `gorm:"type:text" json:"address"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
