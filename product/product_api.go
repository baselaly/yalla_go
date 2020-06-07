package product

import (
	"net/http"
	"yalla_go/request"

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
	validationErrors := request.ValidatePostProductRequest(c)

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrors})
		return
	}

	err := api.ProductService.CreateProduct(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
