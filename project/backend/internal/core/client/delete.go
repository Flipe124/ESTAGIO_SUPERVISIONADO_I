package client

import (
	"encoding/json"
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		DELETE
//	@Description	Deactivate many or all clients.
//	@Tags			client
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			clients	query		[]int	false	"Client ID's."
//	@Success		204		{string}	string	"No Content"
//	@Failure		400		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/client [delete]
func delete(ctx *gin.Context) {

	var err error

	ids := ctx.Query("clients")
	if ids != "" {
		var parsedIds []int
		if err = json.Unmarshal([]byte("["+ids+"]"), &parsedIds); err != nil {
			api.LogReturn(
				ctx,
				http.StatusBadRequest,
				"invalid ids",
				err.Error(),
			)
			return
		}
		err = db.Instance.Delete(&models.Client{}, &parsedIds).Error
	} else {
		err = db.Instance.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Client{}).Error
	}

	if err != nil {
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

// Swagger:
//
//	@Summary		DELETE
//	@Description	Deactivate a single client.
//	@Tags			client
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			client	path		int		true	"Client ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/client/{client} [delete]
func deleteClient(ctx *gin.Context) {

	ID := ctx.Param("client")
	if err := db.Instance.Delete(&models.Client{}, &ID).Error; err != nil {
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
