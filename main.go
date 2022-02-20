package main

import (
	"log"
	"os"

	"github.com/baleghsefat/go-simple-ecommerce/controllers"
	"github.com/baleghsefat/go-simple-ecommerce/database"
	"github.com/baleghsefat/go-simple-ecommerce/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addcart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
