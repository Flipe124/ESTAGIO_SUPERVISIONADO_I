package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User is the struct for ORM operations.
type User struct {
	ID        *uint          `gorm:"primaryKey"`
	Name      string         `gorm:"default:unnamed"`
	Username  *string        `gorm:"unique;not null"`
	Email     *string        `gorm:"unique;not null"`
	Password  *string        `gorm:"unique;not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// HashPassword hash the password of entity.
func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(*user.Password),
		bcrypt.DefaultCost,
	)
	*user.Password = string(hashedPassword)
	return err
}

// IsValidPassword validate the passing password.
func (user *User) IsValidPassword(passwordToCompare string) bool {
	// return bcrypt.CompareHashAndPassword([]byte(passwordToCompare), []byte(testUser.Password)) == nil
	err := bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(passwordToCompare),
	)
	return err == nil
}

// UserList is the struct to bind list GET requests.
type UserList struct {
	ID       *uint   `json:"id"`
	Name     string  `json:"name"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Role     byte    `json:"role"`
}

// UserCreate is the struct to bind create POST requests.
type UserCreate struct {
	Name     string  `json:"name" binding:"omitempty,phrase"`
	Username *string `json:"username" binding:"required,username"`
	Email    *string `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required,min=8,max=32"`
	Role     byte    `json:"role" binding:"omitempty,numeric"`
}

// UserUpdate is the struct to bind update PATCH requests.
type UserUpdate struct {
	Name     string `json:"name,omitempty" binding:"omitempty,phrase"`
	Username string `json:"username,omitempty" binding:"omitempty,username"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Role     byte   `json:"role,omitempty" binding:"omitempty,numeric"`
}
