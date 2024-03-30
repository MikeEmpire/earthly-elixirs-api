package products

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"
)

func ArchiveProduct(context *gin.Context) {
	// Set the Stripe API key from environment variable
	stripe.Key = os.Getenv("KEY")

	// Get product ID and price ID from URL params
	productId := context.Param("product_id")
	priceId := context.Param("price_id")

	// Check if product ID or price ID is missing
	if productId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing product id"})
		return
	}
	if priceId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing price id"})
		return
	}

	// Set parameters for updating price and product
	productParams := &stripe.ProductParams{}
	priceParams := &stripe.PriceParams{}

	// Set 'active' flag to false for archiving
	active := false
	priceParams.Active = &active
	productParams.Active = &active

	// Update the price first
	_, priceErr := price.Update(priceId, priceParams)
	if priceErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to archive price", "error": priceErr.Error()})
		return
	}

	// Update the product after the price is archived
	_, productErr := product.Update(productId, productParams)
	if productErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to archive product", "error": productErr.Error()})
		return
	}

	// Return success message if both price and product are archived successfully
	context.JSON(http.StatusOK, gin.H{"message": "Successfully archived product"})
}
