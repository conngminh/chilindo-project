package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model  `json:"-"`
	Id          int    `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	UserName    string `json:"username" gorm:"unique"`
	Password    string `json:"password,omitempty" gorm:"type:nvarchar(100)"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	Language    string `json:"language"`
	Role        string `json:"role" gorm:"default:user"` //0 is user, 1 is admin
}

func (u *User) HashPassword(password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("Fail to hash password in model user", err)
		return err
	}
	u.Password = string(passHash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		log.Println("error: Check Password error in model user", err)
		return err
	}
	return nil
}
