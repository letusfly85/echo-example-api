package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

// Base Model's definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type Account struct {
	gorm.Model
	Name string
	Hasher string
	Salt string
	Password string
	FirstName string
	LastName string
	Email string
	AvatarUrl string
	Activated bool
}