package account

import (
	"errors"
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		GET
//	@Description	Get a single account from ID.
//	@Tags			account
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			account	path		int		true	"Account ID."
//	@Success		200		{object}	models.AccountList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/account/{account} [get]
func get(ctx *gin.Context) {

	var account *models.Account

	accountList := &models.AccountList{}

	if err := db.Tx.First(&account, ctx.Param("account")).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "account not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(account, accountList)

	ctx.JSON(http.StatusOK, accountList)

}
