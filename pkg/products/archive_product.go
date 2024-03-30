package products

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func ArchiveProduct(context *gin.Context) {
	// Set the Stripe API key from environment variable
	stripe.Key = os.Getenv("KEY")

	// Get product ID from params
	productId := context.Param("product_id")

	// Check if product ID is missing
	if productId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing product id"})
		return
	}

	// Set parameters for updating price and product
	productParams := &stripe.ProductParams{}

	// Set 'active' flag to false for archiving
	active := false
	productParams.Active = &active

	// Update the product
	_, productErr := product.Update(productId, productParams)
	if productErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to archive product", "error": productErr.Error()})
		return
	}

	// Return success message if product is archived
	context.JSON(http.StatusOK, gin.H{"message": "Successfully archived product"})
}
