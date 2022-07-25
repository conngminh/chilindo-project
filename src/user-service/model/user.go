package model

import "gorm.io/gorm"

type User struct {
	gorm.Model  `gorm:"-"`
	Id          int    `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	UserName    string `json:"username" gorm:"unique"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	Language    string `json:"language"`
	Role        string `json:"role" gorm:"default:user"`
}
