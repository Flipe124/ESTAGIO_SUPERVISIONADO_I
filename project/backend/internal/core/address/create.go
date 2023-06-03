package address

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
//	@Description	Create a new address.
//	@Tags			address
//	@Accept			json
//	@Produce		json
//	@Param			TOKEN	header		string					true	"Bearer token."
//	@Param			JSON	body		models.AddressCreate	true	"Json request."
//	@Success		201		{object}	models.AddressList
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/address [post]
func create(ctx *gin.Context) {

	var addressCreate *models.AddressCreate

	address := &models.Address{}
	addressList := &models.AddressList{}

	if err := ctx.ShouldBindJSON(&addressCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	structure.Assign(addressCreate, address)

	if err := db.Instance.Create(&address).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}
	structure.Assign(address, addressList)

	ctx.JSON(http.StatusCreated, addressList)

}
