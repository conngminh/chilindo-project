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
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)
	authRoute := route.NewAuthRoute(authController, r)
	authRoute.SetRoute()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userRoute := route.NewUserRoute(userController, r)
	userRoute.SetRoute()

	if err := r.Run(":1011"); err != nil {
		log.Println("Open port is fail!")
		return
	}
}

func Route() *gin.Engine {
	route := gin.Default()
	return route
}
