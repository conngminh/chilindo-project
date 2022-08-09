package route

import (
	"github.com/gin-gonic/gin"
	"practice/pkg/middleware"
	"practice/src/user-service/controller"
)

type IUserRoute interface {
	SetRoute()
}

type UserRoute struct {
	userController controller.IUserController
	route          *gin.Engine
}

func NewUserRoute(userController controller.IUserController, route *gin.Engine) *UserRoute {
	return &UserRoute{userController: userController, route: route}
}

func (u UserRoute) SetRoute() {
	apiAddress := u.route.Group("/api/admin/address").Use(middleware.IsAuth())
	{
		apiAddress.POST("/create", u.userController.CreateAddress)
	}
}
