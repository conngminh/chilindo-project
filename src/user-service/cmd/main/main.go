package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"practice/src/user-service/controller"
	"practice/src/user-service/database"
	"practice/src/user-service/repository"
	"practice/src/user-service/route"
	"practice/src/user-service/service"
)

func main() {
	db := database.GetDB()
	r := Route()
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	authRoute := route.NewAuthRoute(authController, r)
	authRoute.SetRoute()

	if err := r.Run(":1011"); err != nil {
		log.Println("Open port is fail!")
		return
	}
}

func Route() *gin.Engine {
	route := gin.Default()
	return route
}
