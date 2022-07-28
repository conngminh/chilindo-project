package service

import (
	"log"
	"practice/src/user-service/model"
	"practice/src/user-service/repository"
)

type IUserService interface {
	CreateAddress(address *model.Address) (*model.Address, error)
}

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u UserService) CreateAddress(address *model.Address) (*model.Address, error) {
	address, err := u.userRepo.CreateAddress(address)
	if err != nil {
		log.Println("error to create address in package service")
		return nil, err
	}
	return address, nil
}
