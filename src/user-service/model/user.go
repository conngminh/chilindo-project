package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model  `json:"-"`
	Id          int    `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	Language    string `json:"language"`
	Role        string `json:"role" gorm:"default:admin"`
}

func (u *User) HashPassword(password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("Fail to hash password in model admin", err)
		return err
	}
	u.Password = string(passHash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		log.Println("error: Check Password error in model admin", err)
		return err
	}
	return nil
}
