package address

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
//	@Description	Update already existing address.
//	@Tags			address
//	@Accept			json
//	@Param			TOKEN	header		string					true	"Bearer token."
//	@Param			JSON	body		models.AddressUpdate	true	"Json request."
//	@Param			address	path		int						true	"Address ID."
//	@Success		204		{string}	string					"No Content"
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/address/{address} [patch]
func update(ctx *gin.Context) {

	var addressUpdate *models.AddressUpdate

	ID := ctx.Param("address")
	if err := ctx.ShouldBindJSON(&addressUpdate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}

	if err := db.Instance.Model(&models.Address{}).Where("id", &ID).Updates(&addressUpdate).Error; err != nil {
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
