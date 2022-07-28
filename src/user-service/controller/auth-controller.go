package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"practice/src/user-service/dto"
	"practice/src/user-service/model"
	"practice/src/user-service/service"
	"practice/src/user-service/token"
)

type IAuthController interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type AuthController struct {
	AuthService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (a *AuthController) SignUp(c *gin.Context) {
	var userBody *model.User
	if err := c.ShouldBindJSON(&userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to sign up in controller sign up",
		})
		log.Println("Signup:Error ShouldBindJson in package controller", err.Error())
		return
	}

	user, err := a.AuthService.CreateUser(userBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to sign up duplicate entry",
		})
		log.Println("Signup: Error in package controller")
		return
	}
	c.JSONP(http.StatusOK, user)
}

func (a *AuthController) SignIn(c *gin.Context) {
	var user *dto.SignInDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"sign-in": "Error to should bind json in package controller",
		})
		log.Println("sign-in: Error to should bind json in package controller", err)
		return
	}
	userLogin, errLogin := a.AuthService.GetUserByEmailAndPassword(user)
	if errLogin != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to login",
		})
		log.Println("sign-in: error to login in package controller", errLogin)
		c.Abort()
		return
	}
	tokenString, errToken := token.GenerateJWT(userLogin.Email, userLogin.Username, userLogin.Id, userLogin.Role)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error to generate token in package sign-in controller",
		})
		log.Println("sign-in: error to generateJWT in package controller", errToken)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
