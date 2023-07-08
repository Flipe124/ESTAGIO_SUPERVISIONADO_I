package category

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	category := api.V2.Group("/category", middlewares.Auth)
	{
		category.GET(
			"/:category",
			get,
		)
		category.GET(
			"/",
			list,
		)
		category.POST(
			"/",
			create,
		)
		category.PATCH(
			"/:category",
			update,
		)
		category.DELETE(
			"/:category",
			delete,
		)
	}
}
