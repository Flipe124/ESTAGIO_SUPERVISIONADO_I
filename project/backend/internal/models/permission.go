package models

// Permission is the struct for ORM operations.
type Permission struct {
	ID   *uint   `gorm:"primaryKey"`
	Code *byte   `gorm:"unique;not null"`
	Name *string `gorm:"unique;not null"`
}
