package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
)

type IUserRepository interface {
	CreateUser(dto *dto.SignUpDTO) (*model.User, error)
	GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u UserRepository) CreateUser(dto *dto.SignUpDTO) (*model.User, error) {
	if err := dto.User.HashPassword(dto.User.Password); err != nil {
		log.Println("Sign up:error to hash password in package repository")
		return nil, err
	}
	record := u.Db.Create(&dto.User)
	if record.Error != nil {
		log.Println("sign up: error to sign up in package repository")
		return nil, record.Error
	}
	return dto.User, nil
}

func (u UserRepository) GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error) {
	email := dto.Email
	password := dto.Password
	var user *model.User
	record := u.Db.Where("email =  ?", email).Find(&user)
	if record.Error != nil {
		log.Println("sign-in: error to find user sign-in in package repository ", record.Error)
		return nil, record.Error
	}
	fmt.Println("aaaaaaa--", user.Email, password, email)
	fmt.Println("bbbbbbb--", password)
	if err := user.CheckPassword(password); err != nil {
		log.Println("sign-in: error to check password in package repository ", err)
		return nil, err
	}
	return user, nil
	//email := dto.Email
	//password := dto.Password
	//var user *model.User
	//result := u.Db.Where("email = ?", email).Find(&user)
	//if result.Error != nil {
	//	log.Println("GetUserByUsername: Error find username in package repository", result.Error)
	//	return nil, result.Error
	//}
	////Compose
	//if err := user.CheckPassword(password); err != nil {
	//	log.Println("GetUserByUsername2: Error in check password package repository", err)
	//	return nil, err
	//}
	//return user, nil
}
