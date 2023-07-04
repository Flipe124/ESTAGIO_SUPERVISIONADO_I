package models

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
