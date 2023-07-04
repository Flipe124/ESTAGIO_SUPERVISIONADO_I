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
//	@Description	Deactivate many or all user.
//	@Tags			user
//	@Param			Token	header		string	true	"Bearer token."
//	@Param			users	query		[]int	false	"User ID's."
//	@Success		204		{string}	string	"No Content"
//	@Failure		400		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/user [delete]
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
		err = db.Instance.Delete(&models.User{}, &parsedIds).Error
	} else {
		err = db.Instance.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.User{}).Error
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
//	@Description	Deactivate a single user.
//	@Tags			user
//	@Param			Token	header		string	true	"Bearer token."
//	@Param			user	path		int		true	"User ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/user/{user} [delete]
func deleteUser(ctx *gin.Context) {

	ID := ctx.Param("user")
	if err := db.Instance.Delete(&models.User{}, &ID).Error; err != nil {
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
