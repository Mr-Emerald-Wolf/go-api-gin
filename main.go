package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mr-emerald-wolf/go-api/controllers"
	"github.com/mr-emerald-wolf/go-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main() {
	r := gin.Default()
	r.POST("/create", controllers.CreateProduct)
	r.GET("/findAll", controllers.GetAllProducts)
	r.GET("/find/:id", controllers.FindProduct)

	r.Run()
}

