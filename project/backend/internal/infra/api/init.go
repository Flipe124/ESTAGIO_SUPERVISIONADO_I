package api

import (
	"backend/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

// Router is the singleton instance of api router.
var Router = gin.Default()

// V2 is the singleton instance of v2 route api group.
var V2 = Router.Group(
	"/api/v2",
	middlewares.CORS,
)

func init() {
	V2.OPTIONS("/*any")
}
