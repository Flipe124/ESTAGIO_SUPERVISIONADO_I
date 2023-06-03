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
//	@Summary		LIST
//	@Description	List all addresses.
//	@Tags			address
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.AddressList
//	@Failure		500		{object}	models.HTTP
//	@Router			/address [get]
func list(ctx *gin.Context) {

	var addresses []*models.Address

	if err := db.Instance.Find(&addresses).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	addressesList := make([]*models.AddressList, len(addresses))
	for index, address := range addresses {
		addressesList[index] = &models.AddressList{}
		structure.Assign(address, addressesList[index])
	}

	ctx.JSON(http.StatusOK, addressesList)

}

// Swagger:
//
//	@Summary		LIST
//	@Description	Lists all addresses for a client.
//	@Tags			address
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			client	path		int		true	"Client ID."
//	@Success		200		{array}		models.AddressList
//	@Failure		500		{object}	models.HTTP
//	@Router			/address/client/{client} [get]
func listClient(ctx *gin.Context) {

	var addresses []*models.Address

	clientID := ctx.Param("client")
	if err := db.Instance.Where("client_id", &clientID).Find(&addresses).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	addressesList := make([]*models.AddressList, len(addresses))
	for index, address := range addresses {
		addressesList[index] = &models.AddressList{}
		structure.Assign(address, addressesList[index])
	}

	ctx.JSON(http.StatusOK, addressesList)

}
