package user

import (
	"encoding/json"
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		DELETE
//	@Description	Deactivate many or all finance.
//	@Tags			finance
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			finances	query		[]int	false	"Finance ID's."
//	@Success		204			{string}	string	"No Content"
//	@Failure		400			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/finance [delete]
//	@Deprecated
func delete(ctx *gin.Context) {

	var err error

	ids := ctx.Query("users")
	if ids != "" {
		var parsedIds []int
		if err = json.Unmarshal([]byte("["+ids+"]"), &parsedIds); err != nil {
			api.LogReturn(
				ctx,
				http.StatusBadRequest,
				"invalid ids",
				err.Error(),
			)
			return
		}
		err = db.Tx.Delete(&models.User{}, &parsedIds).Error
	} else {
		err = db.Tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.User{}).Error
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

	ctx.Status(http.StatusNoContent)

}

// Swagger:
//
//	@Summary		DELETE
//	@Description	Deactivate a single finance.
//	@Tags			finance
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			finance	path		int		true	"Finance ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/finance/{finance} [delete]
//	@Deprecated
func deleteUser(ctx *gin.Context) {

	ID := ctx.Param("user")
	if err := db.Tx.Delete(&models.User{}, &ID).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ctx.Status(http.StatusNoContent)

}
