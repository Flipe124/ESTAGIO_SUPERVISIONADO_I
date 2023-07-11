package transaction

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	transaction := api.V2.Group("/transaction", middlewares.Auth)
	{
		transaction.GET(
			"/",
			list,
		)
		transaction.GET(
			"/accounts/",
			listAccounts,
		)
		transaction.GET(
			"/:transaction/",
			get,
		)
		transaction.GET(
			"/:transaction/accounts/",
			getAccounts,
		)
		transaction.POST(
			"/",
			create,
		)
	}
}
