package controller

import (
	"earthly-elixirs-api/pkg/prices"

	"github.com/gin-gonic/gin"
)

func RegisterPricesRoutes(r *gin.Engine) {
	routes := r.Group("/prices")
	routes.POST("/", prices.CreatePrice)
}
