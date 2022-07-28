package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"practice/pkg/config"
	"practice/src/user-service/model"
	"practice/src/user-service/service"
)

type IUserController interface {
	CreateAddress(c *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) CreateAddress(c *gin.Context) {
	userId, ok := c.Get(config.UserId)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "error to get userID in package controller",
		})
		log.Println("error to get userID in package controller")
		c.Abort()
		return
	}
	var address *model.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error to bind json in package controller",
		})
		log.Println("error to bind json in package controller", err)
		c.Abort()
		return
	}
	address.UserId = userId.(int)
	addressUser, err := u.userService.CreateAddress(address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error to create address in package controller",
		})
		log.Println("error to create address in package controller")
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, addressUser)
}
