package finance

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
//	@Description	Delete the finance.
//	@Tags			finance
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			finance	path		int		true	"Finance ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/finance/{finance} [delete]
func delete(ctx *gin.Context) {

	if err := db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Finance{}, ctx.Param("finance")).Error; err != nil {
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
