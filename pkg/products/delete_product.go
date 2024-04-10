package products

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func DeleteProduct(context *gin.Context) {

	// Set the Stripe API key from environment variable
	stripe.Key = os.Getenv("KEY")

	// Get product ID from params
	productId := context.Param("product_id")

	// Check if product ID is missing
	if productId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing product id"})
		return
	}
	params := &stripe.ProductParams{}

	result, err := product.Del(productId, params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"result": result, "productId": productId, "message": "Successfully deleted product!"})
}
