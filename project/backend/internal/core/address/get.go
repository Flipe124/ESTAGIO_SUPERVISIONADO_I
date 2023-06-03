package address

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
//	@Description	Get a single address from ID.
//	@Tags			address
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			address	path		int		true	"Address ID."
//	@Success		200		{object}	models.AddressList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/address/{address} [get]
func get(ctx *gin.Context) {

	var address *models.Address

	addressList := &models.AddressList{}
	ID := ctx.Param("address")

	if err := db.Instance.First(&address, &ID).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "address not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(address, addressList)

	ctx.JSON(http.StatusOK, addressList)

}

// Swagger:
//
//	@Summary		GET
//	@Description	Get a single address for a single client from your ID's.
//	@Tags			address
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			address	path		int		true	"Address ID."
//	@Param			client	path		int		true	"Client ID."
//	@Success		200		{object}	models.AddressList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/address/{address}/client/{client} [get]
func getClient(ctx *gin.Context) {

	var address *models.Address

	addressList := &models.AddressList{}
	ID := ctx.Param("address")
	clientID := ctx.Param("client")

	if err := db.Instance.Where("client_id", &clientID).First(&address, &ID).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "address not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return
	}
	structure.Assign(address, addressList)

	ctx.JSON(http.StatusOK, addressList)

}
