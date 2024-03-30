package prices

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
)

type UpdatePriceInput struct {
	PriceId    string `json:"price_id"`
	UnitAmount int64  `json:"unit_amount"`
}

func UpdatePrice(context *gin.Context) {
	stripe.Key = os.Getenv("KEY")
	var input UpdatePriceInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unitAmount := input.UnitAmount

	params := &stripe.PriceParams{}

	params.UnitAmount = &unitAmount

	result, err := price.Update(input.PriceId, params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to archive price", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated price", "result": result})

}
