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
//	@Param			JSON	body		models.TransactionCreate	true	"Json request."
//	@Success		201		{object}	models.TransactionList
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/transaction [post]
func create(ctx *gin.Context) {

	var transactionCreate *models.TransactionCreate

	transaction := &models.Transaction{}
	transactionList := &models.TransactionList{}

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
	structure.Assign(transaction, transactionList)
	db.Tx.Table("accounts").
		Select("name").
		Where("id", &transactionList.EmitterID).
		Scan(&transactionList.EmitterName)
	db.Tx.Table("accounts").
		Select("name").
		Where("id", &transactionList.BeneficiaryID).
		Scan(&transactionList.BeneficiaryName)

	ctx.JSON(http.StatusCreated, transactionList)

}
