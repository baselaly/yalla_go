//+build wireinject

package app

import (
	"database/sql"
	"yalla_go/user"

	"github.com/google/wire"
)

func InitUserApi(db *sql.DB) user.API {
	wire.Build(user.ProvideUserRepository, user.ProvideUserService, user.ProvideUserAPI)
	return user.API{}
}
