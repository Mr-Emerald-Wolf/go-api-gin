package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mr-emerald-wolf/go-api/controllers"
)

func ProductRoutes(superRoute *gin.RouterGroup) {
	// Simple group: product
	product := superRoute.Group("/product")
	{
		product.POST("/create", controllers.CreateProduct)
		product.GET("/findAll", controllers.FindAllProducts)
		product.GET("/find/:id", controllers.FindProduct)
		product.PUT("/update/:id", controllers.UpdateProduct)
		product.DELETE("/delete/:id", controllers.DeleteProduct)
	}
}
