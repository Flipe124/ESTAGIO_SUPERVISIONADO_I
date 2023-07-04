package user

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
//	@Description	List all available status to use.
//	@Tags			status
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.StatusList
//	@Failure		500		{object}	models.HTTP
//	@Router			/status [get]
func list(ctx *gin.Context) {

	var statuses []*models.Status

	if err := db.Tx.Find(&statuses).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if len(statuses) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	statusesList := make([]*models.StatusList, len(statuses))
	for index, status := range statuses {
		statusesList[index] = &models.StatusList{}
		structure.Assign(status, statusesList[index])
	}

	ctx.JSON(http.StatusOK, statusesList)

}
