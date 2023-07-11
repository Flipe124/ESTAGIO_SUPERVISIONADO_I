package finance

import (
	"net/http"

	"backend/internal/behavior/status"
	"backend/internal/behavior/typet"
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
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/finance/{finance} [patch]
func update(ctx *gin.Context) {

	var (
		financeUpdate *models.FinanceUpdate
		finance       *models.Finance
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

	err = db.Tx.Model(&models.Finance{}).Where("id = ? AND user_id = ?", ctx.Param("finance"), ctx.GetUint("id")).Updates(structure.Map(&financeUpdate)).Error
	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	db.Tx.First(&finance, ctx.Param("finance"))
	if byte(status.Completed) == *finance.StatusCode {
		var balance float64
		db.Tx.Table("accounts").Select("balance").Where("id", &finance.AccountID).Scan(&balance)
		if byte(typet.Input) == *finance.TypeCode {
			balance += *finance.Value
		} else if byte(typet.Output) == *finance.TypeCode {
			balance -= *finance.Value
		} else {
			api.Return(
				ctx,
				http.StatusBadRequest,
				"wrong account type",
			)
		}
		db.Tx.Model(&models.Account{}).Where("id", &finance.AccountID).Update("balance", &balance)
	}

	ctx.Status(http.StatusNoContent)

}
