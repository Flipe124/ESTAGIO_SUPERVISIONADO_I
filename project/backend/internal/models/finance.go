package models

import (
	"time"

	"gorm.io/gorm"
)

// Finance is the struct for ORM operations.
type Finance struct {
	ID *uint `gorm:"primaryKey"`

	

	CreatedAt *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
	// Name      *string         `gorm:"default:unnamed"`
	// Username  *string         `gorm:"unique;not null"`
	// Email     *string         `gorm:"unique;not null"`
	// Password  *string         `gorm:"unique;not null"`
}

// FinanceList is the struct to bind list GET requests.
type FinanceList struct {
	// ID       *uint   `json:"id,omitempty"`
	// Name     *string `json:"name,omitempty"`
	// Username *string `json:"username,omitempty"`
	// Email    *string `json:"email,omitempty"`
	// Role     *byte   `json:"role,omitempty"`
}

// FinanceCreate is the struct to bind create POST requests.
type FinanceCreate struct {
	// Name     *string `json:"name,omitempty" binding:"omitempty,phrase"`
	// Username *string `json:"username" binding:"required,username"`
	// Email    *string `json:"email" binding:"required,email"`
	// Role     *byte   `json:"role,omitempty" binding:"omitempty,numeric"`
	// Password *string `json:"password" binding:"required,min=8,max=32"`
}

// FinanceUpdate is the struct to bind update PATCH requests.
type FinanceUpdate struct {
	// Name     *string `json:"name,omitempty" binding:"omitempty,phrase"`
	// Username *string `json:"username,omitempty" binding:"omitempty,username"`
	// Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	// Role     *byte   `json:"role,omitempty" binding:"omitempty,numeric"`
}
