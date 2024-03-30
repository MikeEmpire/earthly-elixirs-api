package controller

import (
	"earthly-elixirs-api/pkg/products"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	routes := r.Group("/products")
	routes.GET("", products.GetProducts)
	routes.GET("/:product_id", products.GetProductById)
	routes.POST("", products.CreateProduct)
	routes.PUT("", products.UpdateProduct)
	routes.PUT("/archive/:product_id/:price_id", products.ArchiveProduct)
}
