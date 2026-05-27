package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Popup struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`
	Photo     string         `gorm:"type:text;not null" json:"photo"`
	IsRead    bool           `gorm:"type:boolean;default:false;not null" json:"is_read"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Popup) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
