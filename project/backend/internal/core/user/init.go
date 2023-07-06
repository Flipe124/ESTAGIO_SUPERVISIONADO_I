package user

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	user := api.V2.Group("/user")
	{
		user.GET(
			"/:user",
			middlewares.Auth,
			get,
		)
		user.POST(
			"/",
			create,
		)
		user.PATCH(
			"/:user",
			middlewares.Auth,
			update,
		)
		user.DELETE(
			"/:user",
			middlewares.Auth,
			delete,
		)
	}
}
