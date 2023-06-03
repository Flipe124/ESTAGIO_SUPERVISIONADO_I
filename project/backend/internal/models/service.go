package models

import (
	"time"

	"gorm.io/gorm"
)

// Service is the struct for ORM operations.
type Service struct {
	ID        *uint          `gorm:"primaryKey"`
	Service   *string        `gorm:"unique,not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ServiceList is the struct to bind list GET requests.
type ServiceList struct {
	ID      *uint   `json:"id"`
	Service *string `json:"service"`
}

// ServiceCreate is the struct to bind create POST requests.
type ServiceCreate struct {
	Service *string `json:"service" binding:"required,phrase"`
}

// ServiceUpdate is the struct to bind update PATCH requests.
type ServiceUpdate struct {
	Service string `json:"service,omitempty" binding:"omitempty,phrase"`
}
