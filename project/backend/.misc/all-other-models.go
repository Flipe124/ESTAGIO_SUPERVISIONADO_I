package backend

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// --------------- ADDRESS ---------------

// Address is the struct for ORM operations.
type Address struct {
	ID        *uint           `gorm:"primaryKey"`
	Street    *string         `gorm:"not null"`
	Number    *uint           `gorm:"not null"`
	District  *string         `gorm:"not null"`
	City      *string         `gorm:"default:Cianorte"`
	State     *string         `gorm:"default:PR"`
	ClientID  *uint           `gorm:"not null"`
	CreatedAt *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

// AddressList is the struct to bind list GET requests.
type AddressList struct {
	ID       *uint   `json:"id,omitempty"`
	Street   *string `json:"street,omitempty"`
	Number   *uint   `json:"number,omitempty"`
	District *string `json:"district,omitempty"`
	City     *string `json:"city,omitempty"`
	State    *string `json:"state,omitempty"`
	ClientID *uint   `json:"client_id,omitempty"`
}

// AddressCreate is the struct to bind create POST requests.
type AddressCreate struct {
	Street   *string `json:"street" binding:"required,phnum"`
	Number   *uint   `json:"number" binding:"required,numeric"`
	District *string `json:"district" binding:"required,phnum"`
	City     *string `json:"city" binding:"omitempty,phrase"`
	State    *string `json:"state" binding:"omitempty,alpha,len=2"`
	ClientID *uint   `json:"client_id" bindind:"omitempty,numeric"`
}

// AddressUpdate is the struct to bind update PATCH requests.
type AddressUpdate struct {
	Street   *string `json:"street,omitempty" binding:"omitempty,phnum"`
	Number   *uint   `json:"number,omitempty" binding:"omitempty,numeric"`
	District *string `json:"district,omitempty" binding:"omitempty,phnum"`
	City     *string `json:"city,omitempty" binding:"omitempty,phrase"`
	State    *string `json:"state,omitempty" binding:"omitempty,alpha,len=2"`
}

// --------------- AUTH ---------------

// Auth is the struct to manipule authentication operations.
type Auth struct {
	Username *string `json:"username,omitempty" binding:"omitempty,username"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Password *string `json:"password" binding:"required,min=8,max=32"`
}

// --------------- TOKEN ---------------

// Token is the to representate model of token request.
type Token struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}

// --------------- CLIENT ---------------

/*
	The "Email" field can be NULL, however, it has no default value.
	In this case, its type is pointed so that the field's value is "nil" and is set to NULL at insert time.
	If the field was simply the primitive type, it would gain the default value if it did not come at the
	time of binding and in the database it would be an empty string ("") and not NULL.
*/

// Client is the struct for ORM operations.
type Client struct {
	ID        *uint           `gorm:"primaryKey"`
	Name      *string         `gorm:"unique;not null"`
	Document  *string         `gorm:"unique;not null"`
	Phone     *int            `gorm:"unique;not null"`
	Email     *string         `gorm:"unique"`
	CreatedAt *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
	Addresses []*Address
}

// ClientList is the struct to bind list GET requests.
type ClientList struct {
	ID        *uint          `json:"id,omitempty"`
	Name      *string        `json:"name,omitempty"`
	Document  *string        `json:"document,omitempty"`
	Phone     *int           `json:"phone,omitempty"`
	Email     *string        `json:"email,omitempty"`
	Addresses []*AddressList `json:"addresses,omitempty"`
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
	Name     *string `json:"name,omitempty" binding:"omitempty,phnum"`
	Document *string `json:"document,omitempty" binding:"omitempty,numeric,len=9|len=11|len=14"`
	Phone    *int    `json:"phone,omitempty" binding:"omitempty,numeric,length=10|length=11"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
}

// --------------- HTTP ---------------

// HTTP is the struc to representate a error json return in api requests.
type HTTP struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// --------------- ORDER ---------------

// Order is the struct for ORM operations.
type Order struct {
	ID          *uint           `gorm:"primaryKey"`
	State       *byte           `gorm:"not null;default:0"`
	UserID      *uint           `gorm:"not null"`
	ClientID    *uint           `gorm:"not null"`
	ServiceID   *uint           `gorm:"not null"`
	Description *string         `gorm:"default:No description."`
	DateTime    *time.Time      `gorm:"not null;type:datetime"`
	CreatedAt   *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
	Status      *Status         `gorm:"foreignKey:State;references:Code"`
	User        *User
	Client      *Client
	Service     *Service
}

// OrderList is the struct to bind list GET requests.
type OrderList struct {
	ID          *uint        `json:"id,omitempty"`
	State       *byte        `json:"state,omitempty"`
	Description *string      `json:"description,omitempty"`
	DateTime    *time.Time   `json:"date_time,omitempty"`
	User        *UserList    `json:"user,omitempty"`
	Client      *ClientList  `json:"client,omitempty"`
	Service     *ServiceList `json:"service,omitempty"`
}

// OrderCreate is the struct to bind create POST requests.
type OrderCreate struct {
	State       *byte   `json:"state" binding:"omitempty,numeric"`
	UserID      *uint   `json:"user_id" binding:"required,numeric"`
	ClientID    *uint   `json:"client_id" binding:"required,numeric"`
	ServiceID   *uint   `json:"service_id" binding:"required,numeric"`
	Description *string `json:"description" binding:"-"`
	DateTime    *any    `json:"date_time" binding:"required,datetime"`
}

// OrderUpdate is the struct to bind update PATCH requests.
type OrderUpdate struct {
	State       *byte   `json:"state,omitempty" binding:"omitempty,numeric"`
	UserID      *uint   `json:"user_id,omitempty" binding:"omitempty,numeric"`
	ClientID    *uint   `json:"client_id,omitempty" binding:"omitempty,numeric"`
	ServiceID   *uint   `json:"service_id,omitempty" binding:"omitempty,numeric"`
	Description *string `json:"description,omitempty" binding:"-"`
	DateTime    *any    `json:"date_time,omitempty" binding:"omitempty,datetime"`
}

// --------------- PERMISSION ---------------

// Permission is the struct for ORM operations.
type Permission struct {
	ID   *uint   `gorm:"primaryKey"`
	Code *byte   `gorm:"unique;not null"`
	Name *string `gorm:"unique;not null"`
}

// PermissionList is the struct to bind list GET requests.
type PermissionList struct {
	ID   *uint   `json:"id,omitempty"`
	Code *byte   `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// --------------- SERVICE ---------------

// Service is the struct for ORM operations.
type Service struct {
	ID        *uint           `gorm:"primaryKey"`
	Name      *string         `gorm:"unique;not null"`
	CreatedAt *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

// ServiceList is the struct to bind list GET requests.
type ServiceList struct {
	ID   *uint   `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ServiceCreate is the struct to bind create POST requests.
type ServiceCreate struct {
	Name *string `json:"name" binding:"required,phrase"`
}

// ServiceUpdate is the struct to bind update PATCH requests.
type ServiceUpdate struct {
	Name *string `json:"name" binding:"required,phrase"`
}

// --------------- STATUS ---------------

// Status is the struct for ORM operations.
type Status struct {
	ID   *uint   `gorm:"primaryKey"`
	Code *byte   `gorm:"unique;not null"`
	Name *string `gorm:"unique;not null"`
}

// StatusList is the struct to bind list GET requests.
type StatusList struct {
	ID   *uint   `json:"id,omitempty"`
	Code *byte   `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// --------------- USER ---------------

// User is the struct for ORM operations.
type User struct {
	ID         *uint           `gorm:"primaryKey"`
	Name       *string         `gorm:"default:unnamed"`
	Username   *string         `gorm:"unique;not null"`
	Email      *string         `gorm:"unique;not null"`
	Role       *byte           `gorm:"default:0"`
	Password   *string         `gorm:"unique;not null"`
	CreatedAt  *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  *gorm.DeletedAt `gorm:"index"`
	Permission *Permission     `gorm:"foreignKey:Role;references:Code"`
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
	ID       *uint   `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Role     *byte   `json:"role,omitempty"`
}

// UserCreate is the struct to bind create POST requests.
type UserCreate struct {
	Name     *string `json:"name,omitempty" binding:"omitempty,phrase"`
	Username *string `json:"username" binding:"required,username"`
	Email    *string `json:"email" binding:"required,email"`
	Role     *byte   `json:"role,omitempty" binding:"omitempty,numeric"`
	Password *string `json:"password" binding:"required,min=8,max=32"`
}

// UserUpdate is the struct to bind update PATCH requests.
type UserUpdate struct {
	Name     *string `json:"name,omitempty" binding:"omitempty,phrase"`
	Username *string `json:"username,omitempty" binding:"omitempty,username"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Role     *byte   `json:"role,omitempty" binding:"omitempty,numeric"`
}
