package finance

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
//	@Description	List all finances.
//	@Tags			finance
//	@Produce		json
//	@Param			TOKEN		header		string		true	"Bearer token."
//	@Param			account_id	query		uint		false	"Finance Account ID."
//	@Param			type_code	query		byte		false	"Finance Type CODE."
//	@Param			status_code	query		byte		false	"Finance Satus CODE."
//	@Param			category_id	query		uint		false	"Finance Category ID."
//	@Param			value		query		float64		false	"Finance value."
//	@Param			description	query		string		false	"Finance description."
//	@Param			date_time	query		time.Time	false	"Finance date and time."
//	@Param			finances	query		[]int		false	"Finance ID's."
//	@Success		200			{array}		models.FinanceList
//	@Success		204			{string}	string	"No Content"
//	@Failure		500			{object}	models.HTTP
//	@Router			/finance [get]
func list(ctx *gin.Context) {

	var finances []*models.Finance

	tx := db.Tx

	if "" != ctx.Query("finances") {
		tx = tx.Where(strings.Split(ctx.Query("finances"), ","))
	}
	if query, values, paramsExists := query.Make(ctx, &models.FinanceList{}, "ID"); paramsExists {
		tx = tx.Where(query, values...)
	}
	if err := tx.Where("user_id", ctx.GetUint("id")).Find(&finances).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if len(finances) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	financesList := make([]*models.FinanceList, len(finances))
	for index, finance := range finances {
		financesList[index] = &models.FinanceList{}
		structure.Assign(finance, financesList[index])
	}

	ctx.JSON(http.StatusOK, financesList)

}
