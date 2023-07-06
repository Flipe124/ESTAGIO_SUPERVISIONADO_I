package user

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/query"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		LIST
//	@Description	List all finances.
//	@Tags			finance
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			name		query		string	false	"Finance name."
//	@Param			financename	query		string	false	"Finance financename."
//	@Param			email		query		string	false	"Finance email."
//	@Param			role		query		byte	false	"Finance role."
//	@Param			inactives	query		bool	false	"Bring the inactive ones."
//	@Success		200			{array}		models.FinanceList
//	@Failure		500			{object}	models.HTTP
//	@Router			/finance [get]
func list(ctx *gin.Context) {

	var (
		users []*models.User
		err   error
	)

	if ctx.Query("inactives") != "true" {
		if query, values, paramsExists := query.Make(ctx, &models.UserList{}, "ID"); !paramsExists {
			err = db.Tx.Find(&users).Error
		} else {
			err = db.Tx.Where(query, values...).Find(&users).Error
		}
	} else {
		err = db.Tx.Unscoped().Where("deleted_at IS NOT NULL").Find(&users).Error
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

	if len(users) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	usersList := make([]*models.UserList, len(users))
	for index, user := range users {
		usersList[index] = &models.UserList{}
		structure.Assign(user, usersList[index])
	}

	ctx.JSON(http.StatusOK, usersList)

}
