package service

import (
	"log"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
	"practice/src/user-service/repository"
)

type IAuthService interface {
	CreateUser(dto *dto.SignUpDTO) (*model.User, error)
	GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error)
}

type AuthService struct {
	UserRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (a AuthService) CreateUser(dto *dto.SignUpDTO) (*model.User, error) {
	user, err := a.UserRepository.CreateUser(dto)
	if err != nil {
		log.Println("Sign up: error to sign up in package service", err)
		return nil, err
	}
	return user, nil
}

func (a AuthService) GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error) {
	user, err := a.UserRepository.GetUserByEmailAndPassword(dto)
	if err != nil {
		log.Println("sign-in: error to sign-in in package service", err)
		return nil, err
	}
	return user, nil
}
