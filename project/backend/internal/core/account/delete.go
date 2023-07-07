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
//	@Description	Delete the account.
//	@Tags			account
//	@Param			Token	header		string	true	"Bearer token."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/account [delete]
func delete(ctx *gin.Context) {

	id, exists := ctx.Get("id")
	if !exists {
		api.Return(
			ctx,
			http.StatusBadRequest,
			"missing user id",
		)
		return
	}
	if err := db.Tx.Delete(&models.User{}, &id).Error; err != nil {
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
