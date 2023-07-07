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
//	@Description	Create a new account.
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			JSON	body		models.AccountCreate	true	"Json request."
//	@Success		201		{object}	models.AccountList
//	@Failure		409		{object}	models.HTTP
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/account [post]
func create(ctx *gin.Context) {

	var accountCreate *models.AccountCreate

	account := &models.Account{}
	accountList := &models.AccountList{}

	id, exists := ctx.Get("id")
	if !exists {
		api.Return(
			ctx,
			http.StatusBadRequest,
			"missing user id",
		)
		return
	}
	if err := ctx.ShouldBindJSON(&accountCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	structure.Assign(accountCreate, account)
	account.UserID = id.(*uint)

	if err := db.Tx.Create(&account).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if regex.Grep(`(?i)duplicate entry`, err.Error()) {
			code = http.StatusConflict
			message = "name already exists"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(account, accountList)

	ctx.JSON(http.StatusCreated, accountList)

}
