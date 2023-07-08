package models

import (
	"time"

	"gorm.io/gorm"
)

// Account is the struct for ORM operations.
type Account struct {
	ID        *uint           `gorm:"primaryKey"`
	UserID    *uint           `gorm:"not null"`
	Name      *string         `gorm:"not null"`
	Balance   *float64        `gorm:"not null"`
	CreatedAt *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

// AccountList is the struct to bind list GET requests.
type AccountList struct {
	ID      *uint    `json:"id,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
}

// AccountCreate is the struct to bind create POST requests.
type AccountCreate struct {
	Name    *string  `json:"name" binding:"required,phrase"`
	Balance *float64 `json:"balance" binding:"required,gte=0"`
}

// AccountUpdate is the struct to bind update PATCH requests.
type AccountUpdate struct {
	Name    *string  `json:"name,omitempty" binding:"omitempty,phrase"`
	Balance *float64 `json:"balance,omitempty" binding:"omitempty,gte=0"`
}
