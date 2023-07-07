package user

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
//	@Summary		LIST
//	@Description	List all accounts.
//	@Tags			account
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			inactives	query		bool	false	"Bring the inactive ones."
//	@Success		200			{array}		models.AccountList
//	@Failure		500			{object}	models.HTTP
//	@Router			/account [get]
func list(ctx *gin.Context) {

	var (
		accounts []*models.Account
		err      error
	)

	id, exists := ctx.Get("id")
	if !exists {
		api.Return(
			ctx,
			http.StatusBadRequest,
			"missing user id",
		)
		return
	}
	if ctx.Query("inactives") != "true" {
		err = db.Tx.Find(&accounts, &id).Error
	} else {
		err = db.Tx.Unscoped().Where("deleted_at IS NOT NULL").Find(&accounts, &id).Error
	}
	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if len(accounts) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	accountsList := make([]*models.AccountList, len(accounts))
	for index, account := range accounts {
		accountsList[index] = &models.AccountList{}
		structure.Assign(account, accountsList[index])
	}

	ctx.JSON(http.StatusOK, accountsList)

}
