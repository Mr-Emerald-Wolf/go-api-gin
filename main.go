package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mr-emerald-wolf/go-api/controllers"
	"github.com/mr-emerald-wolf/go-api/initializers"
	"github.com/mr-emerald-wolf/go-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main() {
	r := gin.Default()
	router := r.Group("/")
	routes.AddRoutes(router)
	r.Run()
}
