package service

import (
	"log"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
	"practice/src/user-service/repository"
)

type IAuthService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error)
}

type AuthService struct {
	UserRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (a *AuthService) CreateUser(user *model.User) (*model.User, error) {
	userCreate, err := a.UserRepository.CreateUser(user)
	if err != nil {
		log.Println("Sign up: error to sign up in package service", err)
		return nil, err
	}
	return userCreate, nil
}

func (a *AuthService) GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error) {
	user, err := a.UserRepository.GetUserByEmailAndPassword(dto)
	if err != nil {
		log.Println("sign-in: error to sign-in in package service", err)
		return nil, err
	}
	return user, nil
}
