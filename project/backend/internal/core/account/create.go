package account

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
//	@Summary		CREATE
//	@Description	Create a new account.
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			JSON	body		models.AccountCreate	true	"Json request."
//	@Success		201		{object}	models.AccountList
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/account [post]
func create(ctx *gin.Context) {

	var accountCreate *models.AccountCreate

	account := &models.Account{}
	accountList := &models.AccountList{}

	id := ctx.GetUint("id")
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
	account.UserID = &id

	if err := db.Tx.Create(&account).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}
	structure.Assign(account, accountList)

	ctx.JSON(http.StatusCreated, accountList)

}
