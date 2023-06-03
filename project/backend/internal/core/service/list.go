package service

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
//	@Description	List all services.
//	@Tags			service
//	@Produce		json
//	@Param			Token		header		string	true	"Bearer token."
//	@Param			inactives	query		bool	false	"Bring the inactive ones."
//	@Success		200			{array}		models.ServiceList
//	@Failure		500			{object}	models.HTTP
//	@Router			/service [get]
func list(ctx *gin.Context) {

	var (
		services []*models.Service
		err      error
	)

	if ctx.Query("inactives") != "true" {
		err = db.Instance.Find(&services).Error
	} else {
		err = db.Instance.Unscoped().Where("deleted_at IS NOT NULL").Find(&services).Error
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

	servicesList := make([]*models.ServiceList, len(services))
	for index, service := range services {
		servicesList[index] = &models.ServiceList{}
		structure.Assign(service, servicesList[index])
	}

	ctx.JSON(http.StatusOK, servicesList)

}
