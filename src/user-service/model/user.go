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
	Password    string `json:"-" json:"password"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      int    `json:"gender" gorm:"default:0"` // 0 is male, 1 is female, 2 is different
	Language    string `json:"language"`
	Role        int    `json:"role" gorm:"default:0"` //0 is user, 1 is admin
}

func (u *User) HashPassword(password string) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("Fail to hash password in model user")
		return err
	}
	u.Password = string(passHash)
	return nil
}
