package finance

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	finance := api.V2.Group("/finance", middlewares.Auth)
	{
		finance.GET(
			"/",
			list,
		)
		finance.GET(
			"/:finance",
			get,
		)
		finance.POST(
			"/",
			create,
		)
		finance.PATCH(
			"/:finance",
			update,
		)
		finance.DELETE(
			"/:finance",
			delete,
		)
	}
}
