package user

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		LIST
//	@Description	List all users.
//	@Tags			user
//	@Produce		json
//	@Param			Token		header		string	true	"Bearer token."
//	@Param			inactives	query		bool	false	"Bring the inactive ones."
//	@Success		200			{array}		models.UserList
//	@Failure		500			{object}	models.HTTP
//	@Router			/user [get]
func list(ctx *gin.Context) {

	var (
		users []*models.User
		err   error
	)

	if ctx.Query("inactives") != "true" {
		err = db.Instance.Find(&users).Error
	} else {
		err = db.Instance.Unscoped().Where("deleted_at IS NOT NULL").Find(&users).Error
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

	usersList := make([]*models.UserList, len(users))
	for index, user := range users {
		usersList[index] = &models.UserList{}
		structure.Assign(user, usersList[index])
	}

	ctx.JSON(http.StatusOK, usersList)

}
