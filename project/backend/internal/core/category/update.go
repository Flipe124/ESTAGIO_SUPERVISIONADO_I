package category

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		UPDATE
//	@Description	Update the category infos.
//	@Tags			category
//	@Accept			json
//	@Param			TOKEN		header		string					true	"Bearer token."
//	@Param			category	path		int						true	"Category ID."
//	@Param			JSON		body		models.CategoryUpdate	true	"Json request."
//	@Success		204			{string}	string					"No Content"
//	@Failure		422			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/category/{category} [patch]
func update(ctx *gin.Context) {

	var (
		categoryUpdate *models.CategoryUpdate
		err            error
	)

	if err := ctx.ShouldBindJSON(&categoryUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	err = db.Tx.Model(&models.Category{}).Where("id", ctx.Param("category")).Updates(structure.Map(&categoryUpdate)).Error

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
