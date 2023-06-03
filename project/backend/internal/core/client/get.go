package client

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
//	@Description	Get a single client from ID.
//	@Tags			client
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			client	path		int		true	"Client ID."
//	@Success		200		{object}	models.ClientList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/client/{client} [get]
func get(context *gin.Context) {

	var client *models.Client

	clientList := &models.ClientList{}
	ID := context.Param("client")

	if err := db.Instance.First(&client, &ID).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "client not found"
		}

		api.LogReturn(
			context,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(client, clientList, "Address")

	context.JSON(http.StatusOK, clientList)

}

// Swagger:
//
//	@Summary		GET
//	@Description	Get a single client with a single address from your ID's.
//	@Tags			client
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			client	path		int		true	"Client ID."
//	@Param			address	path		int		true	"Address ID."
//	@Success		200		{object}	models.ClientList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/client/{client}/address/{address} [get]
func getAddress(ctx *gin.Context) {

	var client *models.Client

	clientList := &models.ClientList{Addresses: []*models.AddressList{{}}}
	clientID := ctx.Param("client")
	addressID := ctx.Param("address")

	if err := db.Instance.Model(&models.Client{}).Preload("Addresses", &addressID).Find(&client, &clientID).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "client not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return
	}

	structure.Assign(client, clientList, "Address")
	structure.Assign(client.Addresses[0], clientList.Addresses[0])

	ctx.JSON(http.StatusOK, clientList)

}
