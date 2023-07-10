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
	TypeCode    *byte `gorm:"not null"`
	StatusCode  *byte `gorm:"not null"`
	CategoryID  *uint
	Value       *float64 `gorm:"not null"`
	Description *string
	DateTime    *time.Time      `gorm:"type:datetime"`
	CreatedAt   *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
	Type        *Type           `gorm:"foreignKey:TypeCode;references:Code"`
	Status      *Status         `gorm:"foreignKey:StatusCode;references:Code"`
}

// FinanceList is the struct to bind list GET requests.
type FinanceList struct {
	ID          *uint      `json:"id,omitempty"`
	AccountID   *uint      `json:"account_id,omitempty"`
	TypeCode    *byte      `json:"type_code,omitempty"`
	StatusCode  *byte      `json:"status_code,omitempty"`
	CategoryID  *uint      `json:"category_id,omitempty"`
	Value       *float64   `json:"value,omitempty"`
	Description *string    `json:"description,omitempty"`
	DateTime    *time.Time `json:"date_time"`
}

// FinanceCreate is the struct to bind create POST requests.
type FinanceCreate struct {
	AccountID   *uint    `json:"account_id" binding:"required,gte=0"`
	TypeCode    *byte    `json:"type_code" binding:"required,gte=0"`
	StatusCode  *byte    `json:"status_code" binding:"required,gte=0"`
	CategoryID  *uint    `json:"category_id,omitempty" binding:"omitempty,gte=0"`
	Value       *float64 `json:"value" binding:"required,gte=0"`
	Description *string  `json:"description,omitempty" binding:"omitempty"`
	DateTime    *any     `json:"date_time,omitempty" binding:"omitempty,datetime"`
}

// FinanceUpdate is the struct to bind update PATCH requests.
type FinanceUpdate struct {
	AccountID   *uint    `json:"account_id,omitempty" binding:"omitempty,gte=0"`
	TypeCode    *byte    `json:"type_code,omitempty" binding:"omitempty,gte=0"`
	StatusCode  *byte    `json:"status_code,omitempty" binding:"omitempty,gte=0"`
	CategoryID  *uint    `json:"category_id,omitempty" binding:"omitempty,gte=0"`
	Value       *float64 `json:"value,omitempty" binding:"omitempty,gte=0"`
	Description *string  `json:"description,omitempty" binding:"omitempty"`
	DateTime    *any     `json:"date_time,omitempty" binding:"omitempty,datetime"`
}
