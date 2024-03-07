package controller

import (
	"earthly-elixirs-api/pkg/checkout"

	"github.com/gin-gonic/gin"
)

func RegisterCheckoutRoutes(r *gin.Engine) {
	routes := r.Group("/checkout")
	routes.POST("/session", checkout.CreateCheckoutSession)
}
