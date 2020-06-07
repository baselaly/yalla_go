package product

import "github.com/gin-gonic/gin"

// Service struct to product service
type Service struct {
	ProductRepository Repository
}

// ProvideProductService function to provide product service
func ProvideProductService(repo Repository) Service {
	return Service{ProductRepository: repo}
}

// CreateProduct function to create product in service
func (service *Service) CreateProduct(c *gin.Context) {

}
