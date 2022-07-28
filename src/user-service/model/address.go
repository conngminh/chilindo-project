package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model  `json:"-"`
	Id          int    `json:"id" gorm:"primaryKey""`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Province    string `json:"province"`
	District    string `json:"district"`
	SubDistrict string `json:"subDistrict"`
	Address     string `json:"address"`
	TypeAddress string `json:"typeAddress"`
	UserId      int    `json:"userId"`
	User        User   `gorm:"foreignKey:UserId" json:"-"`
}
