package category

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
//	@Description	List all categories.
//	@Tags			category
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			inactives	query		bool	false	"Bring the inactive ones."
//	@Success		200			{array}		models.CategoryList
//	@Failure		500			{object}	models.HTTP
//	@Router			/category [get]
func list(ctx *gin.Context) {

	var (
		categories []*models.Category
		err        error
	)

	err = db.Tx.Find(&categories, ctx.GetUint("id")).Error
	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	if len(categories) < 1 {
		ctx.Status(http.StatusNoContent)
		return
	}

	categoriesList := make([]*models.CategoryList, len(categories))
	for index, category := range categories {
		categoriesList[index] = &models.CategoryList{}
		structure.Assign(category, categoriesList[index])
	}

	ctx.JSON(http.StatusOK, categoriesList)

}
