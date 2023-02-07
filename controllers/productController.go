package controllers

import (
	"log"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/mr-emerald-wolf/go-api/initializers"
	"github.com/mr-emerald-wolf/go-api/models"
)

func FindAllProducts(ctx *gin.Context) {
	// Get all products
	products := []models.Product{}
	result := initializers.DB.Find(&products)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"message": "There was an error",
			"error":   result.Error.Error(),
		})
		log.Fatal(result.Error)
		return
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
			"message": "Product not found",
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

func UpdateProduct(ctx *gin.Context) {

	// Get id from url
	id := ctx.Param("id")

	// Getting data from req body
	body := models.NewProduct{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Finding the product to update
	oldProduct := models.Product{}
	result := initializers.DB.First(&oldProduct, id)

	// If product is not found
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"message": "Product not found",
			"error":   result.Error.Error(),
		})
		return
		// log.Fatal(result.Error)
	}

	// Update product
	result = initializers.DB.Model(&oldProduct).Updates(models.Product{
		Name:        body.Name,
		Seller:      body.Seller,
		Price:       body.Price,
		Description: body.Description})

	// If product is not updated
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"message": "Product not updated",
			"error":   result.Error.Error(),
		})
		return
	}

	// If product is succesfully created
	ctx.JSON(200, gin.H{
		"message": "Product created successfully",
		"data":    oldProduct,
	})
}

func DeleteProduct(ctx *gin.Context)  {
	// Get id from url
	id := ctx.Param("id")

	// Get the product
	result := initializers.DB.Delete(&models.Product{}, id)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"message": "Product not deleted",
			"error":   result.Error.Error(),
		})
		return
		// log.Fatal(result.Error)
	}

	// If product is succesfully deleted
	ctx.JSON(200, gin.H{
		"message": "Product deleted",
	})
}
