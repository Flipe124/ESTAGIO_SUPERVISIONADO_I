package finance

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
//	@Description	List all finances.
//	@Tags			finance
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.FinanceList
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/finance [get]
func list(ctx *gin.Context) {

	var (
		finances []*models.Finance
		err      error
	)

	err = db.Tx.Where("user_id", ctx.GetUint("id")).Find(&finances).Error
	if err != nil {
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
