package service

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
//	@Description	Create a new service.
//	@Tags			service
//	@Accept			json
//	@Produce		json
//	@Param			TOKEN	header		string					true	"Bearer token."
//	@Param			JSON	body		models.ServiceCreate	true	"Json request."
//	@Success		201		{object}	models.ServiceList
//	@Failure		409		{object}	models.HTTP
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/service [post]
func create(ctx *gin.Context) {

	var serviceCreate *models.ServiceCreate

	service := &models.Service{}
	serviceList := &models.ServiceList{}

	if err := ctx.ShouldBindJSON(&serviceCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	structure.Assign(serviceCreate, service)

	if err := db.Instance.Create(&service).Error; err != nil {

		code := http.StatusConflict
		message := http.StatusText(http.StatusInternalServerError)

		if regex.Grep(`(?i)duplicate entry`, err.Error()) {
			db.Instance.Unscoped().Where("service", &service).First(&service)
			if service.DeletedAt.Valid {
				message = "deactivated service"
			} else {
				message = "already exists service"
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
	structure.Assign(service, serviceList)

	ctx.JSON(http.StatusCreated, serviceList)

}
