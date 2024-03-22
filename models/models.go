package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Book model
type Book struct {
	gorm.Model
    ID         uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid

	Title  string `json:"title"`
	Author string `json:"author"`
}
