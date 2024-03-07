package checkout

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type LineItemsInput struct {
	PriceId  string `json:"priceId"`
	Quantity int64  `json:"quantity"`
}

func CreateCheckoutSession(context *gin.Context) {

	var input []LineItemsInput

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

	var checkoutSessionLineItems []*stripe.CheckoutSessionLineItemParams

	for _, lineItem := range input {
		checkoutSessionLineItems = append(checkoutSessionLineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(lineItem.PriceId),
			Quantity: stripe.Int64(lineItem.Quantity),
		})
	}

	params := &stripe.CheckoutSessionParams{LineItems: checkoutSessionLineItems, Mode: stripe.String(string(stripe.CheckoutSessionModePayment))}

	s, err := session.New(params)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"session": s, "message": "Success"})
}
