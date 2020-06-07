//+build wireinject

package app

import (
	"database/sql"
	"yalla_go/product"
	"yalla_go/user"

	"github.com/google/wire"
)

func InitUserApi(db *sql.DB) user.API {
	wire.Build(user.ProvideUserRepository, user.ProvideUserService, user.ProvideUserAPI)
	return user.API{}
}

func InitProductApi(db *sql.DB) product.API {
	wire.Build(product.ProvideProductRepository, product.ProvideProductService, product.ProvideProductAPI)
	return product.API{}
}
