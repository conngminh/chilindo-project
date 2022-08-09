package routes

import (
	"github.com/gin-gonic/gin"
	"practice/src/product-service/controllers"
	controller2 "practice/src/product-service/controllers/admin-rpc"
)

type IProductRoute interface {
	SetRouter()
}

type ProductRoute struct {
	ProductController  controllers.IProductController
	Router             *gin.Engine
	AdminSrvController controller2.IAdminServiceController
}

func NewProductRoute(productController controllers.IProductController, router *gin.Engine, adminSrvController controller2.IAdminServiceController) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router, AdminSrvController: adminSrvController}
}

func (p *ProductRoute) SetRouter() {
	api := p.Router.Group("/api/products")
	{
		api.POST("/create", p.AdminSrvController.CheckIsAdmin(), p.ProductController.CreateProduct)
	}
}
