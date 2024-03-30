package products

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func GetProductById(context *gin.Context) {
	stripe.Key = os.Getenv("KEY")
	productId := context.Param("product_id")
	if productId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing product id"})
		return
	}

	params := &stripe.ProductParams{}
	result, err := product.Get(productId, params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"product": result, "message": "Successfully retrieved product"})
}
