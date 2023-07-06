package models

import (
	"time"

	"gorm.io/gorm"
)

// Finance is the struct for ORM operations.
type Finance struct {
	ID         *uint           `gorm:"primaryKey"`
	UserID     *uint           `gorm:"not null"`
	AccountID  *uint           `gorm:"not null"`
	CategoryID *uint           `gorm:"not null"` // criar duas categorias default, entra e saida (0 e 1?).
	CreatedAt  *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  *gorm.DeletedAt `gorm:"index"`
}

// FinanceList is the struct to bind list GET requests.
type FinanceList struct {
}

// FinanceCreate is the struct to bind create POST requests.
type FinanceCreate struct {
}

// FinanceUpdate is the struct to bind update PATCH requests.
type FinanceUpdate struct {
}
