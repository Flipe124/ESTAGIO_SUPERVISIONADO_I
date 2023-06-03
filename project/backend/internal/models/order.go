package models

import (
	"time"

	"gorm.io/gorm"
)

// Order is the struct for ORM operations.
type Order struct {
	ID          *uint          `gorm:"primaryKey"`
	UserID      *uint          `gorm:"not null"`
	ClientID    *uint          `gorm:"not null"`
	ServiceID   *uint          `gorm:"not null"`
	Description string         `gorm:"type:string"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User        *User
	Client      *Client
	Service     *Service
}

// OrderList is the struct to bind list GET requests.
type OrderList struct {
	ID          *uint        `json:"id"`
	Description string       `json:"description"`
	User        *UserList    `json:"user"`
	Client      *ClientList  `json:"client"`
	Service     *ServiceList `json:"service"`
}

// OrderCreate is the struct to bind create POST requests.
type OrderCreate struct {
	UserID      *uint  `json:"user_id" binding:"required,numeric"`
	ClientID    *uint  `json:"client_id" binding:"required,numeric"`
	ServiceID   *uint  `json:"service_id" binding:"required,numeric"`
	Description string `json:"description" binding:"-"`
}

// OrderUpdate is the struct to bind update PATCH requests.
type OrderUpdate struct {
	UserID      uint   `json:"user_id,omitempty" binding:"omitempty,numeric"`
	ClientID    uint   `json:"client_id,omitempty" binding:"omitempty,numeric"`
	ServiceID   uint   `json:"service_id,omitempty" binding:"omitempty,numeric"`
	Description string `json:"description,omitempty" binding:"-"`
}
