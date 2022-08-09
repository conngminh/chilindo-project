package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	rpcServer "practice/src/user-service/cmd/rpc-server"
	"practice/src/user-service/controller"
	"practice/src/user-service/database"
	"practice/src/user-service/repository"
	"practice/src/user-service/route"
	"practice/src/user-service/service"
)

const (
	addr = ":50051" // Mở rpc server user
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

	go func() {
		if err := r.Run(":1012"); err != nil {
			log.Println("Open port is fail!")
			return
		}
	}()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//set up server mở port
	if err = rpcServer.RunGRPCServer(false, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server admin is running")

}

func Route() *gin.Engine {
	route := gin.Default()
	return route
}
