package transaction

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
//	@Description	List all transactions.
//	@Tags			transaction
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.TransactionList
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/transaction [get]
func list(ctx *gin.Context) {

	var (
		transactions []*models.Transaction
		err          error
	)

	err = db.Tx.Where("user_id", ctx.GetUint("id")).Find(&transactions).Error
	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if len(transactions) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	transactionsList := make([]*models.TransactionList, len(transactions))
	for index, transaction := range transactions {
		transactionsList[index] = &models.TransactionList{}
		structure.Assign(transaction, transactionsList[index], "Emitter", "Beneficiary")
	}

	ctx.JSON(http.StatusOK, transactionsList)

}

// Swagger:
//
//	@Summary		LIST
//	@Description	List all transactions with all related accounts.
//	@Tags			transaction
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.TransactionList
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/transaction/accounts [get]
func listAccounts(ctx *gin.Context) {

	var (
		transactions []*models.Transaction
		err          error
	)

	// err = db.Tx.Joins("Emitter").Joins("Beneficiary").Where("transactions.user_id", ctx.GetUint("id")).Find(&transactions).Error
	err = db.Tx.Preload("Emitter").Preload("Beneficiary").Where("user_id", ctx.GetUint("id")).Find(&transactions).Error
	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if len(transactions) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	transactionsList := make([]*models.TransactionList, len(transactions))
	for index, transaction := range transactions {
		transactionsList[index] = &models.TransactionList{
			Emitter:     &models.AccountList{},
			Beneficiary: &models.AccountList{},
		}
		structure.Assign(transaction, transactionsList[index], "Emitter", "Beneficiary")
		structure.Assign(transaction.Emitter, transactionsList[index].Emitter)
		structure.Assign(transaction.Beneficiary, transactionsList[index].Beneficiary)
	}

	ctx.JSON(http.StatusOK, transactionsList)

}
