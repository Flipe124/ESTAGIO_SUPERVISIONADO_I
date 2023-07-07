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
		structure.Assign(transaction, transactionsList[index])
		db.Tx.Table("accounts").
			Select("name").
			Where("id", &transactionsList[index].EmitterID).
			Scan(&transactionsList[index].EmitterName)
		db.Tx.Table("accounts").
			Select("name").
			Where("id", &transactionsList[index].BeneficiaryID).
			Scan(&transactionsList[index].BeneficiaryName)
	}

	ctx.JSON(http.StatusOK, transactionsList)

}
