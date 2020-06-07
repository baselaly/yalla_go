package product

import "github.com/gin-gonic/gin"

type Service struct {
	ProductRepository Repository
}

func ProvideProductService(repo Repository) Service {
	return Service{ProductRepository: repo}
}

func (service *Service) CreateProduct(c *gin.Context) {

}
