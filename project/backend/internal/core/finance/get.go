package finance

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
//	@Description	Get a single finance from ID.
//	@Tags			finance
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			finance	path		int		true	"Finance ID."
//	@Success		200			{object}	models.FinanceList
//	@Failure		404			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/finance/{finance} [get]
func get(ctx *gin.Context) {

	var finance *models.Finance

	financeList := &models.FinanceList{}

	if err := db.Tx.Where("user_id", ctx.GetUint("id")).First(&finance, ctx.Param("finance")).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "finance not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(finance, financeList)

	ctx.JSON(http.StatusOK, financeList)

}
