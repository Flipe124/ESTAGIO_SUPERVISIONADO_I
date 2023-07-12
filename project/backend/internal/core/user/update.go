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
//	@Description	Update the user infos.
//	@Tags			user
//	@Accept			json
//	@Param			TOKEN	header		string				true	"Bearer token."
//	@Param			JSON	body		models.UserUpdate	true	"Json request."
//	@Success		204		{string}	string				"No Content"
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/user [patch]
func update(ctx *gin.Context) {

	var userUpdate *models.UserUpdate

	if err := ctx.ShouldBindJSON(&userUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}

	if err := db.Tx.Model(&models.User{}).Where("id", ctx.GetUint("id")).Updates(structure.Map(&userUpdate)).Error; err != nil {

		code := http.StatusConflict
		message := http.StatusText(http.StatusInternalServerError)

		if regex.Grep(`(?i)duplicate entry`, err.Error()) {
			message = "already exists"
			switch {
			case regex.Grep(`(?i)username`, err.Error()):
				message = "username " + message
			case regex.Grep(`(?i)email`, err.Error()):
				message = "email " + message
			case regex.Grep(`(?i)password`, err.Error()):
				message = "password " + message
			}
		} else {
			code = http.StatusInternalServerError
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
