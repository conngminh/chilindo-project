package service

import (
	"log"
	"practice/pkg/pb/admin"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
	"practice/src/user-service/repository"
	"practice/src/user-service/token"
	"strings"
)

type IAuthService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error)
	CheckIsAdmin(req *admin.CheckIsAdminRequest) (*admin.CheckIsAdminResponse, error)
}

type AuthService struct {
	AuthRepository repository.IAuthRepository
}

func NewAuthService(authRepository repository.IAuthRepository) *AuthService {
	return &AuthService{AuthRepository: authRepository}
}

func (a *AuthService) CreateUser(user *model.User) (*model.User, error) {
	userCreate, err := a.AuthRepository.CreateUser(user)
	if err != nil {
		log.Println("Sign up: error to sign up in package service", err)
		return nil, err
	}
	return userCreate, nil
}

func (a *AuthService) GetUserByEmailAndPassword(dto *dto.SignInDTO) (*model.User, error) {
	user, err := a.AuthRepository.GetUserByEmailAndPassword(dto)
	if err != nil {
		log.Println("sign-in: error to sign-in in package service", err)
		return nil, err
	}
	return user, nil
}

func (u *AuthService) CheckIsAdmin(req *admin.CheckIsAdminRequest) (*admin.CheckIsAdminResponse, error) {
	isAuth := false
	isAdmin := false
	tokenString := req.Token

	tokenResult := strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := token.ExtractToken(tokenResult)
	if err != nil {
		log.Println("CheckIsAdmin: ", err)
		return nil, err
	}

	isAuth = true
	if claims.Role == "admin" {
		isAdmin = true
	}

	return &admin.CheckIsAdminResponse{
		IsAuth:  isAuth,
		IsAdmin: isAdmin,
	}, nil
}
