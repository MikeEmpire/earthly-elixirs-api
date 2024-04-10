package prices

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"
)

type UpdatePriceInput struct {
	PriceId    string `json:"priceId"`
	ProductId  string `json:"productId"`
	UnitAmount string `json:"unitAmount"`
}

func UpdatePrice(context *gin.Context) {
	stripe.Key = os.Getenv("KEY")
	var input UpdatePriceInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unitAmount, err := strconv.ParseInt(input.UnitAmount, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create new price
	newPriceParams := &stripe.PriceParams{Product: stripe.String(input.ProductId), UnitAmount: stripe.Int64(unitAmount), Currency: stripe.String(string(stripe.CurrencyUSD))}
	result, err := price.New(newPriceParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// update product with new price id
	productParams := &stripe.ProductParams{DefaultPrice: stripe.String(result.ID)}
	_, err = product.Update(input.ProductId, productParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// archive existing price
	existingPriceParams := &stripe.PriceParams{Active: stripe.Bool(false)}
	_, err = price.Update(input.PriceId, existingPriceParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated price", "result": result})

}
