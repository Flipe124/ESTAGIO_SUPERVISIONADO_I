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
//	@Summary		LIST
//	@Description	List all clients only.
//	@Tags			client
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			inactives	query		bool	false	"Bring the inactive ones."
//	@Success		200			{array}		models.ClientList
//	@Failure		500			{object}	models.HTTP
//	@Router			/client [get]
func list(ctx *gin.Context) {

	var (
		clients []*models.Client
		err     error
	)

	if ctx.Query("inactives") != "true" {
		err = db.Instance.Find(&clients).Error
	} else {
		err = db.Instance.Unscoped().Where("deleted_at IS NOT NULL").Find(&clients).Error
	}

	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	clientsList := make([]*models.ClientList, len(clients))
	for indexClient, client := range clients {
		clientsList[indexClient] = &models.ClientList{}
		structure.Assign(client, clientsList[indexClient], "Addresses")
	}

	ctx.JSON(http.StatusOK, clientsList)

}

// Swagger:
//
//	@Summary		LIST
//	@Description	List all clients with all addresses.
//	@Tags			client
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.ClientList
//	@Failure		500		{object}	models.HTTP
//	@Router			/client/addresses [get]
func listAddresses(ctx *gin.Context) {

	var clients []*models.Client

	if err := db.Instance.Model(&models.Client{}).Preload("Addresses").Find(&clients).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	clientsList := make([]*models.ClientList, len(clients))
	for indexClient, client := range clients {
		clientsList[indexClient] = &models.ClientList{Addresses: make([]*models.AddressList, len(client.Addresses))}
		structure.Assign(client, clientsList[indexClient], "Addresses")
		for indexAddress, address := range client.Addresses {
			clientsList[indexClient].Addresses[indexAddress] = &models.AddressList{}
			structure.Assign(address, clientsList[indexClient].Addresses[indexAddress])
		}
	}

	ctx.JSON(http.StatusOK, clientsList)

}

// Swagger:
//
//	@Summary		LIST
//	@Description	List a single client with all addresses.
//	@Tags			client
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			client	path		int		true	"Client ID."
//	@Success		200		{array}		models.ClientList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/client/{client}/addresses [get]
func listAddress(ctx *gin.Context) {

	var client *models.Client

	clientList := &models.ClientList{}
	ID := ctx.Param("client")

	if err := db.Instance.Model(&models.Client{}).Preload("Addresses").Find(&client, &ID).Error; err != nil {

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
	clientList.Addresses = make([]*models.AddressList, len(client.Addresses))
	for index, address := range client.Addresses {
		clientList.Addresses[index] = &models.AddressList{}
		structure.Assign(address, clientList.Addresses[index])
	}

	ctx.JSON(http.StatusOK, clientList)

}
