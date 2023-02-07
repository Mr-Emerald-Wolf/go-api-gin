package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mr-emerald-wolf/go-api/initializers"
	"github.com/mr-emerald-wolf/go-api/models"
)

func CreateProduct(ctx *gin.Context) {
	// Getting data from req body
	body := models.NewProduct{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Creating a new product
	newProduct := models.Product{Name: body.Name, Price: body.Price, Description: body.Description, Seller: body.Seller}
	result := initializers.DB.Create(&newProduct) // pass pointer of data to Create

	if result.Error != nil {
		ctx.Status(500)
	}

	// If product is succesfully created
	ctx.JSON(200, gin.H{
		"message": "Product created successfully",
		"data":    newProduct,
	})
}

func GetAllProducts(ctx *gin.Context) {
	// Get all products
	products := []models.Product{}
	result := initializers.DB.Find(&products)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"message": "There was an error",
			"error":   result.Error.Error(),
		})
		log.Fatal(result.Error)
	}

	// If product is succesfully created
	ctx.JSON(200, gin.H{
		"message": "Products found",
		"data":    products,
	})
}

func FindProduct(ctx *gin.Context) {

	// Get id from url
	id := ctx.Param("id")

	// Get the product
	product := models.Product{}
	result := initializers.DB.First(&product, id)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"message": "There was an error",
			"error":   result.Error.Error(),
		})
		return
		// log.Fatal(result.Error)
	}

	// If product is succesfully created
	ctx.JSON(200, gin.H{
		"message": "Product found",
		"data":    product,
	})
}
