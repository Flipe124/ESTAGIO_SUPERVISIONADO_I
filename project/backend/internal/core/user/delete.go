package user

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		DELETE
//	@Description	Delete the user.
//	@Tags			user
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/user [delete]
func delete(ctx *gin.Context) {

	if err := db.Tx.Unscoped().Delete(&models.User{}, ctx.GetUint("id")).Error; err != nil {
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
