package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int64          `gorm:"not null;uniqueIndex;primaryKe" json:"id"`
	Name 	  string         `gorm:"not null" json:"name"`
	Role 	  string         `gorm:"not null" json:"role"`
	CreatedAt   time.Time      `gorm:"not null;default:now()" json:"created_at,omitempty"`
	UpdatedAt   time.Time      `gorm:"not null;default:now()" json:"updated_at,omitempty"`
	DeleteAt    gorm.DeletedAt `gorm:"null" json:"delete_at,omitempty"`
}
