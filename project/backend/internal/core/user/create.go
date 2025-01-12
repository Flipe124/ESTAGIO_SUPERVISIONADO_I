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
//	@Summary		CREATE
//	@Description	Create a new user.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			JSON	body		models.UserCreate	true	"Json request."
//	@Success		201		{object}	models.UserList
//	@Failure		409		{object}	models.HTTP
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/user [post]
func create(ctx *gin.Context) {

	var userCreate *models.UserCreate

	user := &models.User{}
	userList := &models.UserList{}

	if err := ctx.ShouldBindJSON(&userCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	structure.Assign(userCreate, user)

	if err := user.HashPassword(); err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if err := db.Tx.Create(&user).Error; err != nil {

		code := http.StatusConflict
		message := http.StatusText(http.StatusInternalServerError)

		if regex.Grep(`(?i)duplicate entry`, err.Error()) {
			switch {
			case regex.Grep(`(?i)username`, err.Error()):
				message = "username"
			case regex.Grep(`(?i)email`, err.Error()):
				message = "email"
			case regex.Grep(`(?i)password`, err.Error()):
				message = "password"
			}
			message += " already exists"
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
	structure.Assign(user, userList)

	if err := db.Tx.Exec(postCreate, user.ID).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ctx.JSON(http.StatusCreated, userList)

}
