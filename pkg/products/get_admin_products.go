package products

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

// GetProducts is a handler function for getting a list of products from Stripe
func GetAdminProducts(context *gin.Context) {
	// Set the Stripe API key from environment variable
	stripe.Key = os.Getenv("KEY")

	// Create parameters for listing products with a limit of 3 products
	params := &stripe.ProductListParams{}

	// Initialize an iterator to list products
	iter := product.List(params)

	// Create a slice to store products retrieved from Stripe
	var products []stripe.Product

	// Iterate over the products using the iterator
	for iter.Next() {
		// Append each product to the products slice
		products = append(products, *iter.Product())
	}

	// Check for any error during iteration
	if err := iter.Err(); err != nil {
		// If there is an error, return an internal server error response
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function early
	}

	// If no error occurred, return the list of products as JSON response with HTTP status OK (200)
	context.JSON(http.StatusOK, gin.H{"products": products})
}
