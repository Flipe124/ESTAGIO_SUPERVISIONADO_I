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
//	@Summary		CREATE
//	@Description	Create a new category.
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			TOKEN	header		string					true	"Bearer token."
//	@Param			JSON	body		models.CategoryCreate	true	"Json request."
//	@Success		201		{object}	models.CategoryList
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/category [post]
func create(ctx *gin.Context) {

	var categoryCreate *models.CategoryCreate

	category := &models.Category{}
	categoryList := &models.CategoryList{}

	id := ctx.GetUint("id")
	if err := ctx.ShouldBindJSON(&categoryCreate); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}
	structure.Assign(categoryCreate, category)
	category.UserID = &id

	if err := db.Tx.Create(&category).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}
	structure.Assign(category, categoryList)

	ctx.JSON(http.StatusCreated, categoryList)

}
