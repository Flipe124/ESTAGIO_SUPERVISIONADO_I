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
//	@Summary		CREATE
//	@Description	Create a new transaction.
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			TOKEN	header		string						true	"Bearer token."
//	@Param			JSON	body		models.TransactionCreate	true	"Json request."
//	@Success		201		{object}	models.TransactionList
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/transaction [post]
func create(ctx *gin.Context) {

	var (
		transactionCreate *models.TransactionCreate
		EmitterValue      float64
		BeneficiaryValue  float64
	)

	transactionList := &models.TransactionList{
		Emitter:     &models.AccountList{},
		Beneficiary: &models.AccountList{},
	}
	transaction := &models.Transaction{}

	id := ctx.GetUint("id")
	if err := ctx.ShouldBindJSON(&transactionCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	structure.Assign(transactionCreate, transaction)
	transaction.UserID = &id

	if err := db.Tx.Create(&transaction).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	db.Tx.Table("accounts").Select("balance").Where("id", &transaction.EmitterID).Scan(&EmitterValue)
	EmitterValue -= *transaction.Value
	db.Tx.Model(&models.Account{}).Where("id", &transaction.EmitterID).Update("balance", &EmitterValue)

	db.Tx.Table("accounts").Select("balance").Where("id", &transaction.BeneficiaryID).Scan(&BeneficiaryValue)
	BeneficiaryValue += *transaction.Value
	db.Tx.Model(&models.Account{}).Where("id", &transaction.BeneficiaryID).Update("balance", &BeneficiaryValue)

	structure.Assign(transaction, transactionList, "Emitter", "Beneficiary")
	structure.Assign(transaction.Emitter, transactionList.Emitter)
	structure.Assign(transaction.Beneficiary, transactionList.Beneficiary)

	ctx.JSON(http.StatusCreated, transactionList)

}
