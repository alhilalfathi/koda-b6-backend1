package main

import (
	"koda-b6-backend1/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUser)
	r.GET("/users/:id", handlers.GetUserById)
	r.PATCH("/users/:id", handlers.EditUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/product", handlers.GetProduct)
	r.GET("/product/:id", handlers.GetProductById)
	r.POST("/product", handlers.CreateProduct)

	r.Run("localhost:8989")
}
