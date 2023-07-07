package finance

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
//	@Description	Update the finance infos.
//	@Tags			finance
//	@Accept			json
//	@Param			TOKEN	header		string					true	"Bearer token."
//	@Param			finance	path		int						true	"Finance ID."
//	@Param			JSON	body		models.FinanceUpdate	true	"Json request."
//	@Success		204		{string}	string					"No Content"
//	@Failure		409		{object}	models.HTTP
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/finance/{finance} [patch]
func update(ctx *gin.Context) {

	var (
		financeUpdate *models.FinanceUpdate
		err           error
	)

	if err := ctx.ShouldBindJSON(&financeUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	err = db.Tx.Model(&models.Finance{}).Where("id", ctx.Param("finance")).Updates(structure.Map(&financeUpdate)).Error

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
