package main

import (
	"github.com/gin-gonic/gin"
	"log"
	rpc_client "practice/src/product-service/cmd/rpc-client"
	"practice/src/product-service/controllers"
	controller2 "practice/src/product-service/controllers/admin-rpc"
	"practice/src/product-service/database"
	"practice/src/product-service/repository"
	"practice/src/product-service/routes"
	"practice/src/product-service/services"
)

const (
	//DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	//ginPort              = ":3030"
	addr = ":50051"
)

func main() {
	db := database.GetDB()
	// set up rpc client
	rpcClient := rpc_client.NewRPCClient()
	adminClient := rpcClient.SetUpAdminClient()

	r := router()
	productRepo := repository.NewProductRepository(db)
	productScv := services.NewProductService(productRepo)
	productCtr := controllers.NewProductController(productScv)
	adminSrvCtr := controller2.NewAdminServiceController(adminClient)
	productRoute := routes.NewProductRoute(productCtr, r, adminSrvCtr)
	productRoute.SetRouter()

	if err := r.Run(":1011"); err != nil {
		log.Println("open port is fail!")
		return
	}
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
