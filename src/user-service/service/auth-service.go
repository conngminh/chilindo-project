package service

import (
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
	"practice/src/user-service/repository"
)

type IAuthService interface {
	CreateUser(dto *dto.SignUpSTO) (*model.User, error)
}

type AuthService struct {
	UseRepository *repository.IUserRepository
}

func NewAuthService(useRepository *repository.IUserRepository) *AuthService {
	return &AuthService{UseRepository: useRepository}
}

func (a AuthService) CreateUser(dto *dto.SignUpSTO) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
