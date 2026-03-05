package main

import (
	docs "koda-b6-backend1/docs"
	"koda-b6-backend1/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Backend Apps
// @version 1.0.0
// @description this is basic bakcend apps with DI
// @host localhost:8989
// @BasePath /
func main() {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUser)
	r.GET("/users/:id", handlers.GetUserById)
	r.PATCH("/users/:id", handlers.EditUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.POST("/product", handlers.CreateProduct)
	r.GET("/product", handlers.GetProduct)
	r.GET("/product/:id", handlers.GetProductById)
	r.DELETE("/product/:id", handlers.DeleteProduct)

	docs.SwaggerInfo.BasePath = "/"

	docPath := r.Group("/docs")
	{
		docPath.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run("localhost:8989")
}
