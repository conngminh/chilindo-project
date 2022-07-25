package route

import (
	"github.com/gin-gonic/gin"
	"practice/src/user-service/controller"
)

type IRoute interface {
	setRoute()
}

type Route struct {
	AuthController controller.IAuthController
	Route          *gin.Context
}

func NewRoute(authController controller.IAuthController, route *gin.Context) *Route {
	return &Route{AuthController: authController, Route: route}
}

func (r Route) setRoute() {
	//TODO implement me
	panic("implement me")
}
