package repository

import (
	"gorm.io/gorm"
	"log"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
)

type IAuthRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error)
}

type AuthRepository struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{Db: db}
}

func (u AuthRepository) CreateUser(user *model.User) (*model.User, error) {
	if err := user.HashPassword(user.Password); err != nil {
		log.Println("Sign up:error to hash password in package repository")
		return nil, err
	}
	record := u.Db.Create(&user)
	if record.Error != nil {
		log.Println("sign up: error to sign up in package repository")
		return nil, record.Error
	}
	return user, nil
}

func (u *AuthRepository) GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error) {
	var user *model.User

	record := u.Db.Where("email =  ?", dto.Email).Find(&user)
	if record.Error != nil {
		log.Println("sign-in: error to find user sign-in in package repository ", record.Error)
		return nil, record.Error
	}

	if err := user.CheckPassword(dto.Password); err != nil {
		log.Println("sign-in: error to check password in package repository ", err)
		return nil, err
	}
	return user, nil
}
