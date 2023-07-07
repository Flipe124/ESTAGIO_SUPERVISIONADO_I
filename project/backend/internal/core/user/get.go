package user

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
//	@Description	Get the user infos.
//	@Tags			user
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{object}	models.UserList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/user [get]
func get(ctx *gin.Context) {

	var user *models.User

	userList := &models.UserList{}

	id, exists := ctx.Get("id")
	if !exists {
		api.Return(
			ctx,
			http.StatusBadRequest,
			"missing user id",
		)
		return
	}
	if err := db.Tx.First(&user, &id).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "user not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(user, userList)

	ctx.JSON(http.StatusOK, userList)

}
