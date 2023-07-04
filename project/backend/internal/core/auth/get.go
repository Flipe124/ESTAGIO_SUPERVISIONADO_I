package auth

import (
	"net/http"
	"strings"

	"backend/internal/consts"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt/v2"
)

// Swagger:
//
//	@Summary		PROPERTIES
//	@Description	Get the user id from your token.
//	@Tags			auth
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/auth [get]
func get(ctx *gin.Context) {

	token := strings.TrimPrefix(ctx.GetHeader("Token"), "Bearer ")
	if token == "" {
		api.Return(
			ctx,
			http.StatusBadRequest,
			"request does not contain an access token",
		)
		return
	}

	id, err := jwt.StdGetUserID(token, consts.JWTSECRETKEY)
	api.Return(
		ctx,
		http.StatusBadRequest,
		"invalid field to search",
	)
	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ctx.String(http.StatusOK, "%d", id)

}
