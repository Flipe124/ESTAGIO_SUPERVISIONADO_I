package category

import (
	"net/http"
	"strings"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/query"
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
//	@Param			name		query		string	false	"Category name."
//	@Param			categories	query		[]int	false	"Category ID's."
//	@Success		200			{array}		models.CategoryList
//	@Success		204			{string}	string	"No Content"
//	@Failure		500			{object}	models.HTTP
//	@Router			/category [get]
func list(ctx *gin.Context) {

	var categories []*models.Category

	tx := db.Tx

	if "" != ctx.Query("categories") {
		tx = tx.Where(strings.Split(ctx.Query("categories"), ","))
	}
	if query, values, paramsExists := query.Make(ctx, &models.CategoryList{}, "ID", "Icon"); paramsExists {
		tx = tx.Where(query, values...)
	}
	if err := tx.Where("user_id", ctx.GetUint("id")).Find(&categories).Error; err != nil {
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
