package account

import (
	"net/http"
	"strings"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		DELETE
//	@Description	Delete many or all accounts.
//	@Tags			account
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			accounts	query		[]int	false	"Account ID's."
//	@Success		204			{string}	string	"No Content"
//	@Failure		400			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/account [delete]
func delete(ctx *gin.Context) {

	var result *gorm.DB

	if "" != ctx.Query("accounts") {
		result = db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Account{}, strings.Split(ctx.Query("accounts"), ","))
	} else {
		result = db.Tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Account{})
	}
	if result.Error != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			result.Error.Error(),
		)
		return
	} else if result.RowsAffected < 1 {
		api.Return(
			ctx,
			-1,
			"no removed",
		)
		return
	}

	ctx.Status(http.StatusNoContent)

}

// Swagger:
//
//	@Summary		DELETE
//	@Description	Delete the account.
//	@Tags			account
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			account	path		int		true	"Account ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Failure		000		{string}	string	"No Removed (-1)"
//	@Router			/account/{account} [delete]
func deleteAccount(ctx *gin.Context) {

	if result := db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Account{}, ctx.Param("account")); result.Error != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			result.Error.Error(),
		)
		return
	} else if result.RowsAffected < 1 {
		api.Return(
			ctx,
			-1,
			"no removed",
		)
		return
	}

	ctx.Status(http.StatusNoContent)

}
