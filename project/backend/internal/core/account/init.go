package user

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	user := api.V2.Group("/user")
	{
		user.GET(
			"/",
			middlewares.Auth,
			list,
		)
		user.POST(
			"/",
			create,
		)
		user.PATCH(
			"/",
			middlewares.Auth,
			update,
		)
		user.DELETE(
			"/",
			middlewares.Auth,
			delete,
		)
	}
}
