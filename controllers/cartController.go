package controllers

import "github.com/gin-gonic/gin"

// func GetItemFromCart() gin.HandlerFunc {

// }

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {

		}
	}
}

// func (app *Application) RemoveItem() {

// }

// func (app *Application) BuyFromCart() {

// }

// func (app *Application) InstantBuy() {

// }
