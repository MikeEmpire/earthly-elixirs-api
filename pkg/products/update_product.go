package products

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

type UpdateProductInput struct {
	ProductId   string    `json:"product_id"`
	ProductName *string   `json:"product_name"`
	Category    *string   `json:"category"`
	Images      *[]string `json:"images"`
	TaxCode     *string   `json:"tax_code"`
	Active      *bool     `json:"active"`
}

func UpdateProduct(context *gin.Context) {
	stripe.Key = os.Getenv("KEY")
	var input UpdateProductInput
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := &stripe.ProductParams{}

	if input.Category != nil {
		params.AddMetadata("category", *stripe.String(*input.Category))
	}
	// Update product name if provided
	if input.ProductName != nil {
		// Update product name
		params.Name = input.ProductName
	}

	// Update images if provided
	if input.Images != nil {
		imagePtrs := make([]*string, 0, len(*input.Images))
		for i, image := range *input.Images {
			imagePtr := image
			imagePtrs[i] = &imagePtr
		}
		// Update images
		params.Images = imagePtrs
	}

	// Update tax code if provided
	if input.TaxCode != nil {
		// Update tax code
		params.TaxCode = stripe.String(*input.TaxCode)
	}

	result, err := product.Update(input.ProductId, params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"result": result, "message": "Success!"})
}
