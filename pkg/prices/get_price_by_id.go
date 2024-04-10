package prices

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
)

func GetPriceById(context *gin.Context) {

	stripe.Key = os.Getenv("KEY")
	priceId := context.Param("price_id")
	if priceId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing price id"})
		return
	}

	params := &stripe.PriceParams{}
	result, err := price.Get(priceId, params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"price": result, "message": "Successfully retrieved price"})
}
