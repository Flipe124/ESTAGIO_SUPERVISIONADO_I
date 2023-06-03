package client

import (
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"
	"backend/pkg/utils/regex"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		CREATE
//	@Description	Create a new client.
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			TOKEN	header		string				true	"Bearer token."
//	@Param			JSON	body		models.ClientCreate	true	"Json request."
//	@Success		201		{object}	models.ClientList
//	@Failure		409		{object}	models.HTTP
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/client [post]
func create(ctx *gin.Context) {

	var clientCreate *models.ClientCreate

	client := &models.Client{}
	clientList := &models.ClientList{}

	if err := ctx.ShouldBindJSON(&clientCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}

	structure.Assign(clientCreate, client, "Addresses")
	client.Addresses = make([]*models.Address, len(clientCreate.Addresses))
	for index, address := range clientCreate.Addresses {
		client.Addresses[index] = &models.Address{}
		structure.Assign(address, client.Addresses[index])
	}

	if err := db.Instance.Create(&client).Error; err != nil {

		code := http.StatusConflict
		message := http.StatusText(http.StatusInternalServerError)

		if regex.Grep(`(?i)duplicate entry`, err.Error()) {
			db.Instance.Unscoped().Where("document", &client).First(&client)
			if client.DeletedAt.Valid {
				message = "deactivated client"
			} else {
				message = "already exists"
				switch {
				case regex.Grep(`(?i)name`, err.Error()):
					message = "name " + message
				case regex.Grep(`(?i)document`, err.Error()):
					message = "document " + message
				case regex.Grep(`(?i)phone`, err.Error()):
					message = "phone " + message
				case regex.Grep(`(?i)email`, err.Error()):
					message = "email " + message
				}
			}
		} else {
			code = http.StatusInternalServerError
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}

	structure.Assign(client, clientList, "Addresses")
	clientList.Addresses = make([]*models.AddressList, len(client.Addresses))
	for index, address := range client.Addresses {
		clientList.Addresses[index] = &models.AddressList{}
		structure.Assign(address, clientList.Addresses[index])
	}

	ctx.JSON(http.StatusCreated, clientList)

}
