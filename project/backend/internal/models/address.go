package models

import (
	"time"

	"gorm.io/gorm"
)

// Address is the struct for ORM operations.
type Address struct {
	ID        *uint          `gorm:"primaryKey"`
	Street    *string        `gorm:"not null"`
	Number    *uint          `gorm:"not null"`
	District  *string        `gorm:"not null"`
	City      string         `gorm:"default:Cianorte"`
	State     string         `gorm:"default:PR"`
	ClientID  *uint          `gorm:"not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// AddressList is the struct to bind list GET requests.
type AddressList struct {
	ID       *uint   `json:"id"`
	Street   *string `json:"street"`
	Number   *uint   `json:"number"`
	District *string `json:"district"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	ClientID *uint   `json:"client_id"`
}

// AddressCreate is the struct to bind create POST requests.
type AddressCreate struct {
	Street   *string `json:"street" binding:"required,phnum"`
	Number   *uint   `json:"number" binding:"required,numeric"`
	District *string `json:"district" binding:"required,phnum"`
	City     string  `json:"city" binding:"omitempty,phrase"`
	State    string  `json:"state" binding:"omitempty,alpha,len=2"`
	ClientID *uint   `json:"client_id" bindind:"omitempty,numeric"`
}

// AddressUpdate is the struct to bind update PATCH requests.
type AddressUpdate struct {
	Street   string `json:"street,omitempty" binding:"omitempty,phnum"`
	Number   uint   `json:"number,omitempty" binding:"omitempty,numeric"`
	District string `json:"district,omitempty" binding:"omitempty,phnum"`
	City     string `json:"city,omitempty" binding:"omitempty,phrase"`
	State    string `json:"state,omitempty" binding:"omitempty,alpha,len=2"`
}
