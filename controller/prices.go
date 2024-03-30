package controller

import (
	"earthly-elixirs-api/pkg/prices"

	"github.com/gin-gonic/gin"
)

func RegisterPricesRoutes(r *gin.Engine) {
	routes := r.Group("/prices")
	routes.GET("", prices.GetPrices)
	routes.POST("", prices.CreatePrice)
}
