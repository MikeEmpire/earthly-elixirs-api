package controller

import (
	"earthly-elixirs-api/pkg/products"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	routes := r.Group("/products")
	routes.GET("/", products.GetProducts)
	routes.POST("/", products.CreateProduct)
}
