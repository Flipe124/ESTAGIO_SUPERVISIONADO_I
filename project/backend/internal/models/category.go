package models

import (
	"time"

	"gorm.io/gorm"
)

// Category is the struct for ORM operations.
type Category struct {
	ID        *uint   `gorm:"primaryKey"`
	UserID    *uint   `gorm:"not null"`
	Name      *string `gorm:"not null"`
	Icon      *[]byte
	CreatedAt *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

// CategoryList is the struct to bind list GET requests.
type CategoryList struct {
	ID   *uint   `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Icon *[]byte `json:"icon,omitempty"`
}

// CategoryCreate is the struct to bind create POST requests.
type CategoryCreate struct {
	Name *string `json:"name" binding:"required,alpha"`
	Icon *[]byte `json:"icon,omitempty" binding:"-"`
}

// CategoryUpdate is the struct to bind update PATCH requests.
type CategoryUpdate struct {
	Name *string `json:"name,omitempty" binding:"omitempty,alpha"`
	Icon *[]byte `json:"icon,omitempty" binding:"-"`
}
