package category

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
//	@Description	Delete the category.
//	@Tags			category
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			category	path		int		true	"Category ID."
//	@Success		204			{string}	string	"No Content"
//	@Failure		500			{object}	models.HTTP
//	@Router			/category/{category} [delete]
func delete(ctx *gin.Context) {

	if err := db.Tx.Unscoped().Delete(&models.Category{}, ctx.Param("category")).Error; err != nil {
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
