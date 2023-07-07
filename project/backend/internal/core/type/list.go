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
//	@Description	List all available type to use.
//	@Tags			type
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.TypeList
//	@Failure		500		{object}	models.HTTP
//	@Router			/type [get]
func list(ctx *gin.Context) {

	var types []*models.Type

	if err := db.Tx.Find(&types).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if len(types) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	typesList := make([]*models.StatusList, len(types))
	for index, typet := range types {
		typesList[index] = &models.StatusList{}
		structure.Assign(typet, typesList[index])
	}

	ctx.JSON(http.StatusOK, typesList)

}
