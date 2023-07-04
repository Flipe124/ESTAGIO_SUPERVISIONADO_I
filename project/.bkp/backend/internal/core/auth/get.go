package auth

import (
	"net/http"
	"strings"

	"backend/internal/consts"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt/v2"
)

// Swagger:
//
//	@Summary		ID
//	@Description	Get user ID inside the token (JWT).
//	@Tags			auth
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/auth [get]
func get(ctx *gin.Context) {

	token := strings.TrimPrefix(ctx.GetHeader("Token"), "Bearer ")
	if token == "" {
		ctx.JSON(
			http.StatusBadRequest,
			&models.HTTP{
				Code:  http.StatusBadRequest,
				Error: "request does not contain an access token",
			},
		)
		ctx.Abort()
		return
	}

	ID, err := jwt.MapGetUserID(token, consts.JWTSECRETKEY)
	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ctx.String(http.StatusOK, "%d", ID)

}
