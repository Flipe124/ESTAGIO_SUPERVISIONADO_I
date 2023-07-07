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
//	@Description	Update already existing finance.
//	@Tags			finance
//	@Accept			json
//	@Param			TOKEN		header		string					true	"Bearer token."
//	@Param			finance		path		int						true	"Finance ID."
//	@Param			reactivate	query		bool					false	"Reactivate an inactive finance."
//	@Param			JSON		body		models.FinanceUpdate	true	"Json request."
//	@Success		204			{string}	string					"No Content"
//	@Failure		422			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/finance/{finance} [patch]
//	@Deprecated
func update(ctx *gin.Context) {

	var (
		userUpdate *models.UserUpdate
		err        error
	)

	ID := ctx.Param("user")
	if ctx.Query("reactivate") != "true" {
		if err := ctx.ShouldBindJSON(&userUpdate); err != nil {
			api.LogReturn(
				ctx,
				http.StatusUnprocessableEntity,
				"malformed JSON",
				err.Error(),
			)
			return
		}
		err = db.Tx.Model(&models.User{}).Where("id", &ID).Updates(structure.Map(&userUpdate)).Error
	} else {
		err = db.Tx.Unscoped().Model(&models.User{}).Where("id", &ID).Updates(map[string]any{"deleted_at": nil}).Error
	}

	if err != nil {

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
