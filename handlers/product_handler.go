package handlers

import (
	"fmt"
	"koda-b6-backend1/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllProduct godoc
// @Summary Get All product data
// @Description show all product data
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} models.Products
// @Router /product [get]
func GetProduct(ctx *gin.Context) {
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "List of product",
		Results: models.ProductList,
	})
}

// GetProductById godoc
// @Summary Get product data by id
// @Description show product data by searched id
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
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

// Create Product godoc
// @Summary create new product data
// @Description make new product
// @Tags product
// @Accept json
// @Produce json
// @Param request body models.Product true "fields"
// @Success 200 {object} models.Product
// @Router /product [post]
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

// DeleteProduct godoc
// @Summary Delete product data by id
// @Description remove product data by searched id
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Router /product/{id} [delete]
func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	strId, _ := strconv.Atoi(id)
	for i := range models.ProductList {
		if models.ProductList[i].Id == strId {
			models.ProductList = append(models.ProductList[:i], models.ProductList[i+1:]...)
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "Product deleted",
			})
		}
	}
}
