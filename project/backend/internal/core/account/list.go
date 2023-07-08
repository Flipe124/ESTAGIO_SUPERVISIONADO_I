package account

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
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.AccountList
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/account [get]
func list(ctx *gin.Context) {

	var (
		accounts []*models.Account
		err      error
	)

	err = db.Tx.Where("user_id", ctx.GetUint("id")).Find(&accounts).Error
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
