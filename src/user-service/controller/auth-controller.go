package controller

import (
	"github.com/gin-gonic/gin"
	"practice/src/user-service/service"
)

type IAuthController interface {
	SignUp(c *gin.Context)
}

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (a AuthController) SignUp(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
