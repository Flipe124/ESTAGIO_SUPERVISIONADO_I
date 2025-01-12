package models

import (
	"time"

	"gorm.io/gorm"
)

// Transaction is the struct for ORM operations.
type Transaction struct {
	ID            *uint           `gorm:"primaryKey"`
	UserID        *uint           `gorm:"not null"`
	EmitterID     *uint           `gorm:"not null"`
	BeneficiaryID *uint           `gorm:"not null"`
	Value         *float64        `gorm:"not null"`
	CreatedAt     *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     *gorm.DeletedAt `gorm:"index"`
	Emitter       *Account        `gorm:"foreignKey:EmitterID"`
	Beneficiary   *Account        `gorm:"foreignKey:BeneficiaryID"`
}

// TransactionList is the struct to bind list GET requests.
type TransactionList struct {
	ID            *uint        `json:"id,omitempty"`
	EmitterID     *uint        `json:"emitter_id,omitempty"`
	BeneficiaryID *uint        `json:"beneficiary_id,omitempty"`
	Value         *float64     `json:"value,omitempty"`
	Emitter       *AccountList `json:"emitter,omitempty"`
	Beneficiary   *AccountList `json:"beneficiary,omitempty"`
}

// TransactionCreate is the struct to bind create POST requests.
type TransactionCreate struct {
	EmitterID     *uint    `json:"emitter_id" binding:"required,nefield=BeneficiaryID,gte=0"`
	BeneficiaryID *uint    `json:"beneficiary_id" binding:"required,nefield=EmitterID,gte=0"`
	Value         *float64 `json:"value" binding:"required,gte=0"`
}
