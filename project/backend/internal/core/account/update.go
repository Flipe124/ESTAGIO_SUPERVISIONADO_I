package user

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"
	"backend/pkg/utils/regex"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		UPDATE
//	@Description	Update the account infos.
//	@Tags			account
//	@Accept			json
//	@Param			TOKEN	header		string					true	"Bearer token."
//	@Param			JSON	body		models.AccountUpdate	true	"Json request."
//	@Success		204		{string}	string					"No Content"
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/account [patch]
func update(ctx *gin.Context) {

	var (
		accountUpdate *models.AccountUpdate
		err           error
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
	if err := ctx.ShouldBindJSON(&accountUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	err = db.Tx.Model(&models.User{}).Where("id", &id).Updates(structure.Map(&accountUpdate)).Error

	if err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if regex.Grep(`(?i)duplicate entry`, err.Error()) {
			code = http.StatusConflict
			message = "name already exists"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}

	ctx.Status(http.StatusNoContent)

}
