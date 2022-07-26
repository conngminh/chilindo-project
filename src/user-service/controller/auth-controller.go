package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"practice/src/user-service/dto"
	"practice/src/user-service/service"
)

type IAuthController interface {
	SignUp(c *gin.Context)
}

type AuthController struct {
	AuthService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (a AuthController) SignUp(c *gin.Context) {
	var userBody dto.SignUpDTO
	if err := c.ShouldBindJSON(&userBody.User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to sign up in controller sign up",
		})
		log.Println("Signup:Error ShouldBindJson in package controller", err.Error())
		return
	}
	if err := userBody.User.HashPassword(userBody.User.Password); err != nil {
		log.Println("Sign up:error to hash password in package controller")
		return
	}
	user, err := a.AuthService.CreateUser(&userBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to sign up duplicate entry",
		})
		log.Println("Signup: Error in package controller")
		return
	}
	c.JSONP(http.StatusOK, user)
}
