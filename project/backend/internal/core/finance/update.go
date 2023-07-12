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
		financeOld    *models.Finance
		financeNew    *models.Finance
		balance       float64
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

	db.Tx.First(&financeOld, ctx.Param("finance"))
	if err := db.Tx.Model(&models.Finance{}).Where("id = ? AND user_id = ?", ctx.Param("finance"), ctx.GetUint("id")).Updates(structure.Map(&financeUpdate)).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}
	db.Tx.First(&financeNew, ctx.Param("finance"))

	if *financeOld.AccountID != *financeNew.AccountID {
		if byte(status.Completed) == *financeOld.StatusCode {
			db.Tx.Table("accounts").Select("balance").Where("id", &financeOld.AccountID).Scan(&balance)
			if byte(typet.Input) == *financeOld.TypeCode {
				balance -= *financeOld.Value
			} else if byte(typet.Output) == *financeOld.TypeCode {
				balance += *financeOld.Value
			} else {
				api.Return(
					ctx,
					http.StatusInternalServerError,
					"wrong account type",
				)
			}
			db.Tx.Model(&models.Account{}).Where("id", &financeOld.AccountID).Update("balance", &balance)
		}
		if byte(status.Completed) == *financeNew.StatusCode {
			db.Tx.Table("accounts").Select("balance").Where("id", &financeNew.AccountID).Scan(&balance)
			if byte(typet.Input) == *financeNew.TypeCode {
				balance += *financeNew.Value
			} else if byte(typet.Output) == *financeNew.TypeCode {
				balance -= *financeNew.Value
			} else {
				api.Return(
					ctx,
					http.StatusInternalServerError,
					"wrong account type",
				)
			}
			db.Tx.Model(&models.Account{}).Where("id", &financeNew.AccountID).Update("balance", &balance)
		}
	} else {
		if *financeOld.StatusCode != *financeNew.StatusCode {
			db.Tx.Table("accounts").Select("balance").Where("id", &financeNew.AccountID).Scan(&balance)
			if byte(typet.Input) == *financeNew.TypeCode {
				if byte(status.Completed) == *financeNew.StatusCode {
					balance += *financeNew.Value
				} else if byte(status.Pending) == *financeNew.StatusCode {
					balance -= *financeNew.Value
				} else {
					api.Return(
						ctx,
						http.StatusInternalServerError,
						"wrong status type",
					)
				}
			} else if byte(typet.Output) == *financeNew.TypeCode {
				if byte(status.Completed) == *financeNew.StatusCode {
					balance -= *financeNew.Value
				} else if byte(status.Pending) == *financeNew.StatusCode {
					balance += *financeNew.Value
				} else {
					api.Return(
						ctx,
						http.StatusInternalServerError,
						"wrong status type",
					)
				}
			} else {
				api.Return(
					ctx,
					http.StatusInternalServerError,
					"wrong account type",
				)
			}
			db.Tx.Model(&models.Account{}).Where("id", &financeNew.AccountID).Update("balance", &balance)
		}
	}

	ctx.Status(http.StatusNoContent)

}
