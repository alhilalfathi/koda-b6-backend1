package handlers

import (
	"fmt"
	"koda-b6-backend1/models"

	"github.com/gin-gonic/gin"
)

func GetProduct(ctx *gin.Context) {
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "List of product",
		Results: models.ProductList,
	})
}

func GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	for i := range models.ProductList {
		if fmt.Sprint(models.ProductList[i].Id) == id {
			ctx.JSON(200, models.Response{
				Success: true,
				Message: fmt.Sprintf("Product: %s", models.ProductList[i].ProductName),
				Results: models.ProductList[i],
			})
			return
		}
	}
	ctx.JSON(404, models.Response{
		Success: false,
		Message: "Product not found",
	})
}

func CreateProduct(ctx *gin.Context) {
	var data models.Product

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Create product failed",
		})
		return
	}
	if data.ProductName == "" {
		ctx.JSON(400, models.Response{
			Success: true,
			Message: "Product name cannot blank",
		})
		return
	}

	if data.Price <= 0 {
		ctx.JSON(400, models.Response{
			Success: true,
			Message: "Product price must greater than 0",
		})
		return
	}

	data.Id = models.NextProductId
	models.NextProductId++

	models.ProductList = append(models.ProductList, data)
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Product created successfully",
	})
}
