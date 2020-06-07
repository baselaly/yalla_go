package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API struct define product api
type API struct {
	ProductService Service
}

// ProvideProductAPI function to provide productapi
func ProvideProductAPI(ProductService Service) API {
	return API{ProductService: ProductService}
}

// CreateProduct function to create product
func (api *API) CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
