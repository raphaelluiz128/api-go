package main

import (
	"hello/controller"
	"hello/db"
	"hello/repository"
	"hello/usecase"

	"github.com/gin-gonic/gin"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")

}
