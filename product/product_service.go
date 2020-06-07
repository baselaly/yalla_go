package product

import (
	"yalla_go/jwt"

	"github.com/gin-gonic/gin"
)

// Service struct to product service
type Service struct {
	ProductRepository Repository
}

// ProvideProductService function to provide product service
func ProvideProductService(repo Repository) Service {
	return Service{ProductRepository: repo}
}

// CreateProduct function to create product in service
func (service *Service) CreateProduct(c *gin.Context) error {
	name := c.PostForm("name")
	description := c.PostForm("description")
	userID, err := jwt.VerifyToken(c)

	if err != nil {
		return err
	}

	var product Product

	product.SetName(name)
	product.SetDescription(description)
	product.SetUserID(userID)

	err = service.ProductRepository.Create(product)

	if err != nil {
		return err
	}

	return nil
}
