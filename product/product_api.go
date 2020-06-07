package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	ProductService Service
}

func ProvideProductAPI(ProductService Service) API {
	return API{ProductService: ProductService}
}

func (api *API) CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
