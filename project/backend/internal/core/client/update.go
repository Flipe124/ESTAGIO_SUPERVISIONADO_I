package client

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		UPDATE
//	@Description	Update already existing client.
//	@Tags			client
//	@Accept			json
//	@Param			TOKEN	header		string				true	"Bearer token."
//	@Param			JSON	body		models.ClientUpdate	true	"Json request."
//	@Param			client	path		int					true	"Client ID."
//	@Success		204		{string}	string				"No Content"
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/client/{client} [patch]
func update(ctx *gin.Context) {

	var clientUpdate *models.ClientUpdate

	ID := ctx.Param("client")
	if err := ctx.ShouldBindJSON(&clientUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}

	if err := db.Instance.Model(&models.Client{}).Where("id", &ID).Updates(&clientUpdate).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ctx.Status(http.StatusNoContent)

}
