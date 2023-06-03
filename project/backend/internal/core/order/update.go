package order

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		UPDATE
//	@Description	Update already existing order.
//	@Tags			order
//	@Accept			json
//	@Param			Token	header		string				true	"Bearer token."
//	@Param			order	path		int					true	"Order ID."
//	@Param			JSON	body		models.OrderUpdate	true	"Json request."
//	@Success		204		{string}	string				"No Content"
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/order/{order} [patch]
func update(ctx *gin.Context) {

	var serviceUpdate *models.ServiceUpdate

	ID := ctx.Param("service")
	if err := ctx.ShouldBindJSON(&serviceUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}

	if err := db.Instance.Model(&models.Service{}).Where("id", &ID).Updates(&serviceUpdate).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ctx.Status(http.StatusNoContent)

}
