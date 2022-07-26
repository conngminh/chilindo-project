package repository

import (
	"gorm.io/gorm"
	"log"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
)

type IUserRepository interface {
	CreateUser(dto *dto.SignUpDTO) (*model.User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u UserRepository) CreateUser(dto *dto.SignUpDTO) (*model.User, error) {
	record := u.Db.Create(&dto.User)
	if record.Error != nil {
		log.Println("sign up: error to sign up in package repository")
		return nil, record.Error
	}
	return dto.User, nil
}
