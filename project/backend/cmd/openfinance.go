package main

import (
	// Built-in imports.
	"fmt"

	// Project imports.
	"backend/internal/consts"
	"backend/internal/infra/api"
	"backend/pkg/helpers/logger"

	// Load imports.
	_ "backend/internal/process"
)

// Swagger:
//
//	@title			OpenFinance API
//	@version		0.0.0
//	@description	API for OpenFinance project.
//
//	@contact.name	Rhuan Patriky
//	@contact.url	https://linktr.ee/rhuanpk
//	@contact.email	support@rhuanpk.com
//
//	@host			localhost:9999
//	@BasePath		/api/v0
func main() {
	logger.Log.Fatal(api.Router.Run(fmt.Sprintf(":%d", consts.APIPORT)).Error())
}
