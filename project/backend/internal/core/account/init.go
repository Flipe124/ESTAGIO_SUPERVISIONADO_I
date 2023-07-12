package account

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	account := api.V2.Group("/account", middlewares.Auth)
	{
		account.GET(
			"/",
			list,
		)
		account.GET(
			"/:account",
			get,
		)
		account.POST(
			"/",
			create,
		)
		account.PATCH(
			"/:account",
			update,
		)
		account.DELETE(
			"/",
			delete,
		)
		account.DELETE(
			"/:account",
			deleteAccount,
		)
	}
}
