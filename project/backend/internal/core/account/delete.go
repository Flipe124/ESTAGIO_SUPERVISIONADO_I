package account

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
//	@Description	Delete the account.
//	@Tags			account
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			account	path		int		true	"Account ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/account/{account} [delete]
func delete(ctx *gin.Context) {

	if err := db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Account{}, ctx.Param("account")).Error; err != nil {
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
