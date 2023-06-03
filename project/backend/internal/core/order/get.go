package order

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
//	@Description	Get a single order from ID.
//	@Tags			order
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			order	path		int		true	"Order ID."
//	@Success		200		{object}	models.OrderList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/order/{order} [get]
func get(ctx *gin.Context) {

	var service *models.Service

	serviceList := &models.ServiceList{}
	ID := ctx.Param("service")

	if err := db.Instance.First(&service, &ID).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "service not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(service, serviceList)

	ctx.JSON(http.StatusOK, serviceList)

}
