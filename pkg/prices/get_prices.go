package prices

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
)

// GetPrices is a handler function for getting a list of prices from Stripe
func GetPrices(context *gin.Context) {
	// Set the Stripe API key from environment variable
	stripe.Key = os.Getenv("KEY")

	// Create parameters for listing prices with a limit of 20 prices
	params := &stripe.PriceListParams{}
	params.Limit = stripe.Int64(20)

	// Initialize an iterator to list prices
	iter := price.List(params)

	// Create a slice to store prices retrieved from Stripe
	var prices []stripe.Price

	// Iterate over the prices using the iterator
	for iter.Next() {
		// Append each price to the prices slice
		prices = append(prices, *iter.Price())
	}

	// Check for any error during iteration
	if err := iter.Err(); err != nil {
		// If there is an error, return an internal server error response
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function early
	}

	// If no error occurred, return the list of prices as JSON response with HTTP status OK (200)
	context.JSON(http.StatusOK, gin.H{"prices": prices})
}
