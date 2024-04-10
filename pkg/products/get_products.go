package products

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func GetProducts(context *gin.Context) {
	var query string
	stripe.Key = os.Getenv("KEY")
	// Extract parameters from headers
	category := context.GetHeader("category")
	query = fmt.Sprintf("active:'true' AND metadata['category']:'%s'", category)
	if category == "All" {
		query = "active:'true'"
	}
	params := &stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: query,
		},
	}
	iter := product.Search(params)
	var products []stripe.Product
	for iter.Next() {
		products = append(products, *iter.Product())
	}
	if err := iter.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"products": products})
}
