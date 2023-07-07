package models

import (
	"time"

	"gorm.io/gorm"
)

// Finance is the struct for ORM operations.
type Finance struct {
	ID          *uint `gorm:"primaryKey"`
	UserID      *uint `gorm:"not null"`
	AccountID   *uint `gorm:"not null"`
	TypeID      *uint `gorm:"not null;references:Code"`
	StatusID    *uint `gorm:"not null;references:Code"`
	CategoryID  *uint
	Value       *float64 `gorm:"not null"`
	Description *string
	DateTime    *time.Time      `gorm:"type:datetime"`
	CreatedAt   *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}

// FinanceList is the struct to bind list GET requests.
type FinanceList struct {
	ID          *uint      `json:"id,omitempty"`
	AccountID   *uint      `json:"account_id,omitempty"`
	TypeID      *uint      `json:"type_id,omitempty"`
	CategoryID  *uint      `json:"category_id,omitempty"`
	StatusID    *uint      `json:"status_id,omitempty"`
	Value       *float64   `json:"value,omitempty"`
	Description *string    `json:"description,omitempty"`
	DateTime    *time.Time `json:"datetime"`
}

// FinanceCreate is the struct to bind create POST requests.
type FinanceCreate struct {
	AccountID   *uint      `json:"account_id" binding:"required,gt=0"`
	TypeID      *uint      `json:"type_id" binding:"required,gt=0"`
	CategoryID  *uint      `json:"category_id,omitempty" binding:"omitempty,gt=0"`
	StatusID    *uint      `json:"status_id" binding:"required,gt=0"`
	Value       *float64   `json:"value" binding:"required,gt=0"`
	Description *string    `json:"description,omitempty" binding:"omitempty"`
	DateTime    *time.Time `json:"datetime,omitempty" binding:"omitempty,datetime"`
}

// FinanceUpdate is the struct to bind update PATCH requests.
type FinanceUpdate struct {
	AccountID   *uint      `json:"account_id,omitempty" binding:"omitempty,gt=0"`
	TypeID      *uint      `json:"type_id,omitempty" binding:"omitempty,gt=0"`
	CategoryID  *uint      `json:"category_id,omitempty" binding:"omitempty,gt=0"`
	StatusID    *uint      `json:"status_id,omitempty" binding:"omitempty,gt=0"`
	Value       *float64   `json:"value,omitempty" binding:"omitempty,gt=0"`
	Description *string    `json:"description,omitempty" binding:"omitempty"`
	DateTime    *time.Time `json:"datetime,omitempty" binding:"omitempty,datetime"`
}
