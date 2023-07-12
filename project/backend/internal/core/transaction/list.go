package transaction

import (
	"net/http"
	"strings"

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
//	@Description	List all transactions.
//	@Tags			transaction
//	@Produce		json
//	@Param			TOKEN			header		string	true	"Bearer token."
//	@Param			emitter_id		query		uint	false	"Transaction Emitter ID."
//	@Param			beneficiary_id	query		uint	false	"Transaction Beneficiary ID."
//	@Param			value			query		float64	false	"Transaction Value."
//	@Param			transactions	query		[]int	false	"Transaction ID's."
//	@Success		200				{array}		models.TransactionList
//	@Success		204				{string}	string	"No Content"
//	@Failure		500				{object}	models.HTTP
//	@Router			/transaction [get]
func list(ctx *gin.Context) {

	var (
		transactions []*models.Transaction
		err          error
	)

	tx := db.Tx

	if "" != ctx.Query("transactions") {
		tx = tx.Where(strings.Split(ctx.Query("transactions"), ","))
	}
	if query, values, paramsExists := query.Make(ctx, &models.TransactionList{}, "ID", "Emitter", "Beneficiary"); paramsExists {
		tx = tx.Where(query, values...)
	}
	if err = tx.Where("user_id", ctx.GetUint("id")).Find(&transactions).Error; err != nil {
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
//	@Param			TOKEN			header		string	true	"Bearer token."
//	@Param			emitter_id		query		uint	false	"Transaction Emitter ID."
//	@Param			beneficiary_id	query		uint	false	"Transaction Beneficiary ID."
//	@Param			value			query		float64	false	"Transaction Value."
//	@Param			transactions	query		[]int	false	"Transaction ID's."
//	@Success		200				{array}		models.TransactionList
//	@Success		204				{string}	string	"No Content"
//	@Failure		500				{object}	models.HTTP
//	@Router			/transaction/accounts [get]
func listAccounts(ctx *gin.Context) {

	var (
		transactions []*models.Transaction
		err          error
	)

	tx := db.Tx

	if "" != ctx.Query("transactions") {
		tx = tx.Where(strings.Split(ctx.Query("transactions"), ","))
	}
	if query, values, paramsExists := query.Make(ctx, &models.TransactionList{}, "ID", "Emitter", "Beneficiary"); paramsExists {
		tx = tx.Where(query, values...)
	}
	// err = tx.Joins("Emitter").Joins("Beneficiary").Where("transactions.user_id", ctx.GetUint("id")).Find(&transactions).Error
	if err = tx.Preload("Emitter").Preload("Beneficiary").Where("user_id", ctx.GetUint("id")).Find(&transactions).Error; err != nil {
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
