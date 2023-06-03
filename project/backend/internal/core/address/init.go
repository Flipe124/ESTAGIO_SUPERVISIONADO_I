package address

import (
	"backend/internal/infra/api"
	"backend/internal/permission"
	"backend/pkg/middlewares"
)

func init() {
	address := api.V2.Group("/address", middlewares.Auth)
	{
		address.GET(
			"/",
			middlewares.Require(permission.Employee),
			list,
		)
		address.GET(
			"/client/:client",
			middlewares.Require(permission.Employee),
			listClient,
		)
		address.GET(
			"/:address",
			middlewares.Require(permission.Employee),
			get,
		)
		address.GET(
			"/:address/client/:client",
			middlewares.Require(permission.Employee),
			getClient,
		)
		address.POST(
			"/",
			middlewares.Require(permission.Admin),
			create,
		)
		address.PATCH(
			"/:address",
			middlewares.Require(permission.Admin),
			update,
		)
		address.DELETE(
			"/",
			middlewares.Require(permission.Admin),
			delete,
		)
		address.DELETE(
			"/:address",
			middlewares.Require(permission.Admin),
			deleteAddress,
		)
	}
}
