package products

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

type ProductInput struct {
	ProductName string `json:"productName"`
	UnitAmount  int64  `json:"unitAmount"`
	Recurring   bool   `json:"recurring"`
}

func CreateProduct(context *gin.Context) {
	stripe.Key = os.Getenv("KEY")
	var input ProductInput
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

	params := &stripe.ProductParams{
		Name: stripe.String(input.ProductName),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			UnitAmount: stripe.Int64(input.UnitAmount),
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
		},
	}
	params.AddExpand("default_price")
	result, err := product.New(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success", "result": result})

}
