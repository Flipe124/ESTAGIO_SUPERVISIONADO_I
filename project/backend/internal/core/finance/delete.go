package finance

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
//	@Description	Delete many or all finances.
//	@Tags			finance
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			finances	query		[]int	false	"Finance ID's."
//	@Success		204			{string}	string	"No Content"
//	@Failure		400			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/finance [delete]
func delete(ctx *gin.Context) {

	var result *gorm.DB

	if "" != ctx.Query("finances") {
		result = db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Finance{}, strings.Split(ctx.Query("finances"), ","))
	} else {
		result = db.Tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Finance{})
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
//	@Description	Delete the finance.
//	@Tags			finance
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			finance	path		int		true	"Finance ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Failure		000		{string}	string	"No Removed (-1)"
//	@Router			/finance/{finance} [delete]
func deleteFinance(ctx *gin.Context) {

	if result := db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Finance{}, ctx.Param("finance")); result.Error != nil {
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
