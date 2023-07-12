package transaction

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
//	@Description	Get a single transaction.
//	@Tags			transaction
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			transaction	path		int		true	"Transaction ID."
//	@Success		200			{array}		models.TransactionList
//	@Success		204			{string}	string	"No Content"
//	@Failure		500			{object}	models.HTTP
//	@Router			/transaction/{transaction} [get]
func get(ctx *gin.Context) {

	var transaction *models.Transaction

	transactionList := &models.TransactionList{}

	if err := db.Tx.Where("user_id", ctx.GetUint("id")).First(&transaction, ctx.Param("transaction")).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "transaction not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(transaction, transactionList, "Emitter", "Beneficiary")

	ctx.JSON(http.StatusOK, transactionList)

}

// Swagger:
//
//	@Summary		LIST
//	@Description	Get a single transaction with your accounts.
//	@Tags			transaction
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			transaction	path		int		true	"Transaction ID."
//	@Success		200			{array}		models.TransactionList
//	@Success		204			{string}	string	"No Content"
//	@Failure		500			{object}	models.HTTP
//	@Router			/transaction/{transaction}/accounts [get]
func getAccounts(ctx *gin.Context) {

	var transaction *models.Transaction

	transactionList := &models.TransactionList{
		Emitter:     &models.AccountList{},
		Beneficiary: &models.AccountList{},
	}

	if err := db.Tx.Preload("Emitter").Preload("Beneficiary").Where("user_id", ctx.GetUint("id")).First(&transaction, ctx.Param("transaction")).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "transaction not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(transaction, transactionList, "Emitter", "Beneficiary")
	structure.Assign(transaction.Emitter, transactionList.Emitter)
	structure.Assign(transaction.Beneficiary, transactionList.Beneficiary)

	ctx.JSON(http.StatusOK, transactionList)

}
