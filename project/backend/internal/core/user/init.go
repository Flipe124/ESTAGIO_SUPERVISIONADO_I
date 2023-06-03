package user

import (
	"backend/internal/infra/api"
	"backend/internal/permission"
	"backend/pkg/middlewares"
)

func init() {
	user := api.V2.Group("/user", middlewares.Auth, middlewares.Require(permission.Admin))
	{
		user.GET(
			"/",
			list,
		)
		user.GET(
			"/:user",
			get,
		)
		user.POST(
			"/",
			create,
		)
		user.PATCH(
			"/:user",
			update,
		)
		user.DELETE(
			"/",
			delete,
		)
		user.DELETE(
			"/:user",
			deleteUser,
		)
	}
}
