package client

import (
	"backend/internal/infra/api"
	"backend/internal/permission"
	"backend/pkg/middlewares"
)

func init() {
	client := api.V2.Group("/client", middlewares.Auth, middlewares.Require(permission.Employee))
	{
		client.GET(
			"/",
			middlewares.Require(permission.Employee),
			list,
		)
		client.GET(
			"/addresses",
			middlewares.Require(permission.Employee),
			listAddresses,
		)
		client.GET(
			"/:client/addresses",
			middlewares.Require(permission.Employee),
			listAddress,
		)
		client.GET(
			"/:client",
			middlewares.Require(permission.Employee),
			get,
		)
		client.GET(
			"/:client/address/:address",
			middlewares.Require(permission.Employee),
			getAddress,
		)
		client.POST(
			"/",
			middlewares.Require(permission.Admin),
			create,
		)
		client.PATCH(
			"/:client",
			middlewares.Require(permission.Admin),
			update,
		)
		client.DELETE(
			"/",
			middlewares.Require(permission.Admin),
			delete,
		)
		client.DELETE(
			"/:client",
			middlewares.Require(permission.Admin),
			deleteClient,
		)
	}
}
