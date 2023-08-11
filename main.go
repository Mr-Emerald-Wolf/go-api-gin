package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mr-emerald-wolf/go-api/initializers"
	"github.com/mr-emerald-wolf/go-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.Connect()
}

func main() {
	r := gin.Default()
	router := r.Group("/")
	r.GET("/", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)

	})
	routes.AddRoutes(router)
	r.Run()
}
