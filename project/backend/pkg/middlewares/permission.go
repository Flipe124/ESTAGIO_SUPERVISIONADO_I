package middlewares

import (
	"net/http"

	"backend/internal/consts"
	"backend/internal/permission"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt/v2"
)

// Require is the middleware that validate if the request has permissions for any endpoints.
func Require(minPermission permission.Permission) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, existsToken := ctx.Get("token")
		if !existsToken {
			api.Return(ctx, http.StatusBadRequest, "missing token")
			return
		}

		tokenPermission, err := jwt.MapGetKey("rle", token.(string), consts.JWTSECRETKEY)
		if err != nil {
			api.LogReturn(
				ctx,
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				err.Error(),
			)
			return
		}

		if permission.Permission(tokenPermission.(float64)) < minPermission {
			api.Return(ctx, http.StatusForbidden, "invalid permission")
			return
		}

		ctx.Next()

	}
}
