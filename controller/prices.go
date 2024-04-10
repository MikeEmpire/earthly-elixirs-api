package controller

import (
	"earthly-elixirs-api/pkg/prices"

	"github.com/gin-gonic/gin"
)

func RegisterPricesRoutes(r *gin.Engine) {
	routes := r.Group("/prices")
	routes.GET("", prices.GetPrices)
	routes.GET("/:price_id", prices.GetPriceById)
	routes.POST("", prices.CreatePrice)
	routes.PUT("", prices.UpdatePrice)
}
