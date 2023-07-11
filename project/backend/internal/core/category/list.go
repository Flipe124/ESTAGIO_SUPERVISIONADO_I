package category

import (
	"net/http"

	"backend/internal/core/user"
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
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			name	query		string	false	"Category name."
//	@Success		200		{array}		models.CategoryList
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/category [get]
func list(ctx *gin.Context) {

	var (
		categories []*models.Category
		err        error
	)

	tx := db.Tx

	if query, values, paramsExists := query.Make(ctx, &models.CategoryList{}, "ID", "Icon"); paramsExists {
		tx = tx.Where(query, values...)
	}
	err = tx.Where("user_id", user.SystemUserID).Or("user_id", ctx.GetUint("id")).Find(&categories).Error
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
