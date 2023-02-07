package main

import (
	"github.com/mr-emerald-wolf/go-api/initializers"
	"github.com/mr-emerald-wolf/go-api/models"
)

func init()  {
	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main()  {
	initializers.DB.AutoMigrate(&models.Product{})

}