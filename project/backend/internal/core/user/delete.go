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
//	@Description	Deactivate a single user.
//	@Tags			user
//	@Param			Token	header		string	true	"Bearer token."
//	@Param			user	path		int		true	"User ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/user/{user} [delete]
func delete(ctx *gin.Context) {

	ID := ctx.Param("user")
	if err := db.Tx.Unscoped().Delete(&models.User{}, &ID).Error; err != nil {
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
