package models

import "github.com/jinzhu/gorm"

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
	CreateUserId uint
	CreateServiceId uint
	UpdateUserId uint
	UpdateServiceId uint
	DeleteUserId uint
	DeleteServiceId uint
}