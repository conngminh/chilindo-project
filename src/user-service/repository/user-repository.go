package repository

import (
	"gorm.io/gorm"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
)

type IUserRepository interface {
	CreateUser(dto *dto.SignUpSTO) (*model.User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u UserRepository) CreateUser(dto *dto.SignUpSTO) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
