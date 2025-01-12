package finance

import (
	"net/http"

	"backend/internal/behavior/status"
	"backend/internal/behavior/typet"
	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		CREATE
//	@Description	Create a new finance.
//	@Description	The "datetime" field follows this pattern: "yyyy-mm-dd hh:mm:ss".
//	@Tags			finance
//	@Accept			json
//	@Produce		json
//	@Param			TOKEN	header		string					true	"Bearer token."
//	@Param			JSON	body		models.FinanceCreate	true	"Json request."
//	@Success		201		{object}	models.FinanceList
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/finance [post]
func create(ctx *gin.Context) {

	var financeCreate *models.FinanceCreate

	finance := &models.Finance{}
	financeList := &models.FinanceList{}

	id := ctx.GetUint("id")
	if err := ctx.ShouldBindJSON(&financeCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	structure.Assign(financeCreate, finance)
	finance.UserID = &id

	if err := db.Tx.Create(&finance).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}
	structure.Assign(finance, financeList)

	if byte(status.Completed) == *finance.StatusCode {
		var balance float64
		db.Tx.Table("accounts").Select("balance").Where("id", &finance.AccountID).Scan(&balance)
		if byte(typet.Input) == *finance.TypeCode {
			balance += *finance.Value
		} else if byte(typet.Output) == *finance.TypeCode {
			balance -= *finance.Value
		} else {
			api.Return(
				ctx,
				http.StatusBadRequest,
				"wrong account type",
			)
		}
		db.Tx.Model(&models.Account{}).Where("id", &finance.AccountID).Update("balance", &balance)
	}

	ctx.JSON(http.StatusCreated, financeList)

}
