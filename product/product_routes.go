package product

import (
	"yalla_go/middleware"

	"github.com/gin-gonic/gin"
)

// Routes generate routes for products
func Routes(route *gin.Engine, PRODUCTAPI API) {
	private := route.Group("/product")

	private.Use(middleware.Auth())
	private.POST("/post", PRODUCTAPI.CreateProduct)
}
