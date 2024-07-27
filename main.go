package main

import (
	"log"
	"os"

	"github.com/aungmyatmoe11/GO-Ecommerce/middleware"
	"github.com/aungmyatmoe11/GO-Ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// app := controllers.NewApplication(database.CollectionData(database.Client, "Products"), database.CollectionData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	// ! routes
	routes.UserRoutes(router)
	router.Use(middleware.Authentication()) // ! authentication
	// router.GET("/addtocart", app.AddToCart())
	// router.GET("/removeitem", app.RemoveItem())
	// router.GET("/listcart", controllers.GetItemFromCart())
	// router.POST("/addaddress", controllers.AddAddress())
	// router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	// router.PUT("/editworkaddress", controllers.EditWorkAddress())
	// router.GET("/deleteaddresses", controllers.DeleteAddress())
	// router.GET("/cartcheckout", app.BuyFromCart())
	// router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))
}
