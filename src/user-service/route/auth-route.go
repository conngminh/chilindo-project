package route

import (
	"github.com/gin-gonic/gin"
	"practice/src/user-service/controller"
)

type IAuthRoute interface {
	setRoute()
}

type AuthRoute struct {
	AuthController controller.IAuthController
	Route          *gin.Engine
}

func NewAuthRoute(authController controller.IAuthController, authRoute *gin.Engine) *AuthRoute {
	return &AuthRoute{AuthController: authController, Route: authRoute}
}

func (a AuthRoute) SetRoute() {
	api := a.Route.Group("/api/auth")
	{
		api.POST("/signup", a.AuthController.SignUp)
		api.POST("/signin", a.AuthController.SignIn)
	}
}
