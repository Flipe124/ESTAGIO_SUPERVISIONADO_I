package models

import (
	"time"

	"gorm.io/gorm"
)

/*
	The "Email" field can be NULL, however, it has no default value.
	In this case, its type is pointed so that the field's value is "nil" and is set to NULL at insert time.
	If the field was simply the primitive type, it would gain the default value if it did not come at the
	time of binding and in the database it would be an empty string ("") and not NULL.
*/

// Client is the struct for ORM operations.
type Client struct {
	ID        *uint          `gorm:"primaryKey"`
	Name      *string        `gorm:"unique;not null"`
	Document  *string        `gorm:"unique;not null"`
	Phone     *int           `gorm:"unique;not null"`
	Email     *string        `gorm:"unique"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Addresses []*Address
}

// ClientList is the struct to bind list GET requests.
type ClientList struct {
	ID        *uint          `json:"id"`
	Name      *string        `json:"name"`
	Document  *string        `json:"document"`
	Phone     *int           `json:"phone"`
	Email     *string        `json:"email"`
	Addresses []*AddressList `json:"addresses"`
}

// ClientCreate is the struct to bind create POST requests.
type ClientCreate struct {
	Name      *string          `json:"name" binding:"required,phnum"`
	Document  *string          `json:"document" binding:"required,numeric,len=9|len=11|len=14"`
	Phone     *int             `json:"phone" binding:"required,numeric,length=10|length=11"`
	Email     *string          `json:"email" binding:"omitempty,email"`
	Addresses []*AddressCreate `json:"addresses" binding:"required,dive"`
}

// ClientUpdate is the struct to bind update PATCH requests.
type ClientUpdate struct {
	Name     string `json:"name,omitempty" binding:"omitempty,phnum"`
	Document string `json:"document,omitempty" binding:"omitempty,numeric,len=9|len=11|len=14"`
	Phone    int    `json:"phone,omitempty" binding:"omitempty,numeric,length=10|length=11"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
}
