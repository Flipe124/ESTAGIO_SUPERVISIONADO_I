package user

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		UPDATE
//	@Description	Update already existing user.
//	@Tags			user
//	@Accept			json
//	@Param			Token	header		string				true	"Bearer token."
//	@Param			user	path		int					true	"User ID."
//	@Param			JSON	body		models.UserUpdate	true	"Json request."
//	@Success		204		{string}	string				"No Content"
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/user/{user} [patch]
func update(ctx *gin.Context) {

	var userUpdate *models.UserUpdate

	ID := ctx.Param("user")
	if err := ctx.ShouldBindJSON(&userUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}

	if err := db.Instance.Model(&models.User{}).Where("id", &ID).Updates(&userUpdate).Error; err != nil {
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
