package prices

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
)

type PriceInput struct {
	ProductId  string  `json:"product_id"`
	UnitAmount float64 `json:"unit_amount"`
}

func CreatePrice(context *gin.Context) {
	stripe.Key = os.Getenv("KEY")
	var input PriceInput
	if err := context.ShouldBindJSON(&input); err != nil {

		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage, "error_msg": "Invalid JSON input"})
		return
	}

	params := &stripe.PriceParams{
		Product:    stripe.String(input.ProductId),
		UnitAmount: stripe.Int64(int64(input.UnitAmount)),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
	}

	result, err := price.New(params)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"result": result, "message": "Successfully created price"})
	return
}
