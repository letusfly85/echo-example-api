package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"golang.org/x/crypto/argon2"
	"encoding/hex"
)

// Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type Account struct {
	gorm.Model
	Hasher string
	Salt string
	Password string `gorm:"not null"`
	FirstName string
	LastName string
	Email string `gorm:"not null;unique"`
	AvatarUrl string
	Activated bool
	CreateUserId uint
	CreateServiceId uint
	UpdateUserId uint
	UpdateServiceId uint
	DeleteUserId uint
	DeleteServiceId uint
}

type Accounts []Account

//todo get from configuration
var salt = []byte("sample_salt")

func (a *Account) Create(param Account) (err error) {
	account := Account{}
	account.Hasher = "argon2" //todo change dynamically from configuration

	account.Email = param.Email

	key := argon2.Key([]byte(param.Password), salt, 4, 32*1024, 1, 32)
	account.Password = hex.EncodeToString(key[:])
	account.Salt = hex.EncodeToString(salt[:])

	if err := db.Create(&account).Error; err != nil {
		fmt.Printf("%v\n", err.Error())
		return err
	}

	if err := db.Find(&account).Error; err != nil {
		fmt.Printf("%v\n", err.Error())
		return err
	}

	return nil
}

func (a *Account) Find(param Account) (err error) {
	account := Account{}
	account.Email = param.Email

	key := argon2.Key([]byte(param.Password), salt, 4, 32*1024, 1, 32)
	expectedPassword := hex.EncodeToString(key[:])

	if err := db.Where("email = ?", param.Email).First(&account).Error; err != nil {
		fmt.Printf("%v\n", err.Error())
		return err
	}

	if expectedPassword != account.Password {
		fmt.Printf("%v\n", "invalid password")
		return fmt.Errorf("invalid email or password")
	}

	return nil
}

func (a *Account) Delete(param Account) (err error) {
	fmt.Printf("%v\n", param.Email)

	account := Account{}
	account.Email = param.Email
	if err := db.Unscoped().Delete(&account).Error; err != nil {
		fmt.Printf("%v\n", err.Error())
		return err
	}

	return nil
}

func (as *Accounts) List() (err error) {
	if err := db.Find(&as).Error; err != nil {
		fmt.Printf("%v\n", err.Error())
		return err
	}

	return nil
}
