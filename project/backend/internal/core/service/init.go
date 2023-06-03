package service

import (
	"backend/internal/infra/api"
	"backend/internal/permission"
	"backend/pkg/middlewares"
)

func init() {
	user := api.V2.Group("/service", middlewares.Auth, middlewares.Require(permission.Admin))
	{
		user.GET(
			"/",
			list,
		)
		user.GET(
			"/:service",
			get,
		)
		user.POST(
			"/",
			create,
		)
		user.PATCH(
			"/:service",
			update,
		)
		user.DELETE(
			"/",
			delete,
		)
		user.DELETE(
			"/:service",
			deleteService,
		)
	}
}
