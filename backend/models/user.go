package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Role      string         `gorm:"type:varchar(50);not null" json:"role"`
	StudentID *uuid.UUID     `gorm:"type:uuid" json:"student_id,omitempty"`
	Student   *Student       `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	Avatar    string         `gorm:"type:varchar(255)" json:"avatar"`
	NIP       string         `gorm:"column:nip;type:varchar(50);uniqueIndex" json:"nip,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
