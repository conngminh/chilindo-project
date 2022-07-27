package dto

import "practice/src/user-service/model"

type SignUpDTO struct {
	User *model.User
}

type SignInDTO struct {
	Email    string
	Password string
}
