package address

import (
	"encoding/json"
	"net/http"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		DELETE
//	@Description	Permanently delete many or all address.
//	@Tags			address
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			addresses	query		[]int	false	"Address ID's."
//	@Success		204			{string}	string	"No Content"
//	@Failure		400			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/address [delete]
func delete(ctx *gin.Context) {

	var err error

	ids := ctx.Query("addresses")
	if ids != "" {
		var parsedIds []int
		if err = json.Unmarshal([]byte("["+ids+"]"), &parsedIds); err != nil {
			api.LogReturn(
				ctx,
				http.StatusBadRequest,
				"invalid ids",
				err.Error(),
			)
			return
		}
		err = db.Instance.Delete(&models.Address{}, &parsedIds).Error
	} else {
		err = db.Instance.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Address{}).Error
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

	ctx.Status(http.StatusNoContent)

}

// Swagger:
//
//	@Summary		DELETE
//	@Description	Permanently delete an address.
//	@Tags			address
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			address	path		int		true	"Address ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/address/{address} [delete]
func deleteAddress(ctx *gin.Context) {

	ID := ctx.Param("address")
	if err := db.Instance.Delete(&models.Address{}, &ID).Error; err != nil {
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
