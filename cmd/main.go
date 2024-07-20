package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	usecase "go-api/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConn, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConn)
	ProductUsecases := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecases)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProducts)
	server.POST("/product/:id", ProductController.UpdateProductById)

	server.Run(":8000")
}
