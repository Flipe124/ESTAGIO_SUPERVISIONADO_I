package category

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	category := api.V2.Group("/category", middlewares.Auth)
	{
		category.GET(
			"/default/",
			defaultList,
		)
		category.GET(
			"/default/:default",
			defaultGet,
		)
		category.GET(
			"/",
			list,
		)
		category.GET(
			"/:category",
			get,
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
			"/",
			delete,
		)
		category.DELETE(
			"/:category",
			deleteCategory,
		)
	}
}
