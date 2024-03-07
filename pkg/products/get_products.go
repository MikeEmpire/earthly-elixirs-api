package products

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func GetProducts(context *gin.Context) {

	stripe.Key = os.Getenv("KEY")

	params := &stripe.ProductListParams{}
	result := product.List(params)
	context.JSON(http.StatusOK, gin.H{"products": result})
}
