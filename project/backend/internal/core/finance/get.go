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
//	@Description	Get a single finance from ID.
//	@Tags			finance
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			finance	path		int		true	"Finance ID."
//	@Success		200		{object}	models.FinanceList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/finance/{finance} [get]
func get(ctx *gin.Context) {

	var user *models.User

	userList := &models.UserList{}
	ID := ctx.Param("user")

	if err := db.Tx.First(&user, &ID).Error; err != nil {

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
