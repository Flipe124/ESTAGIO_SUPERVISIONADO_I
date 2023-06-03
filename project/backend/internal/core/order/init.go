package order

import (
	"backend/internal/infra/api"
	"backend/internal/permission"
	"backend/pkg/middlewares"
)

func init() {
	user := api.V2.Group("/order", middlewares.Auth, middlewares.Require(permission.Admin))
	{
		user.GET(
			"/",
			list,
		)
		user.GET(
			"/:order",
			get,
		)
		user.POST(
			"/",
			create,
		)
		user.PATCH(
			"/:order",
			update,
		)
		user.DELETE(
			"/",
			delete,
		)
		user.DELETE(
			"/:order",
			deleteOrder,
		)
	}
}
