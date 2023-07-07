package models

// Type is the struct for ORM operations.
type Type struct {
	ID   *uint   `gorm:"primaryKey"`
	Code *byte   `gorm:"unique;not null"`
	Name *string `gorm:"unique;not null"`
}

// TypeList is the struct to bind list GET requests.
type TypeList struct {
	ID   *uint   `json:"id,omitempty"`
	Code *byte   `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}
