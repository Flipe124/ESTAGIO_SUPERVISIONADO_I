package middlewares

import (
	"net/http"
	"strings"

	"backend/internal/consts"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt/v2"
)

// Auth is the middleware that's validate if request have a token and if its valid.
func Auth(ctx *gin.Context) {

	token := strings.TrimPrefix(ctx.GetHeader("Token"), "Bearer ")
	if token == "" {
		api.Return(ctx, http.StatusBadRequest, "missing token")
		return
	}

	if isValid, err := jwt.StdIsValid(token, consts.JWTSECRETKEY); !isValid {
		api.LogReturn(
			ctx,
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			err.Error(),
		)
		return
	}

	ctx.Set("token", token)
	ctx.Next()

}
