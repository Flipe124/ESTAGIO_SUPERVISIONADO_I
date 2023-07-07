package user

import (
	"backend/internal/infra/api"
	"backend/pkg/middlewares"
)

func init() {
	api.V2.GET(
		"/type/",
		middlewares.Auth,
		list,
	)
}
