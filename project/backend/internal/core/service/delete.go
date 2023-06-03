package service

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
//	@Description	Deactivate many or all services.
//	@Tags			service
//	@Param			Token		header		string	true	"Bearer token."
//	@Param			services	query		[]int	false	"Service ID's."
//	@Success		204			{string}	string	"No Content"
//	@Failure		400			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/service [delete]
func delete(ctx *gin.Context) {

	var err error

	ids := ctx.Query("services")
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
		err = db.Instance.Delete(&models.Service{}, &parsedIds).Error
	} else {
		err = db.Instance.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Service{}).Error
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
//	@Description	Deactivate a single service.
//	@Tags			service
//	@Param			Token	header		string	true	"Bearer token."
//	@Param			service	path		int		true	"Service ID."
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/service/{service} [delete]
func deleteService(ctx *gin.Context) {

	ID := ctx.Param("service")
	if err := db.Instance.Delete(&models.Service{}, &ID).Error; err != nil {
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
