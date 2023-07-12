package category

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
//	@Description	Get a single category from ID.
//	@Tags			category
//	@Produce		json
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			category	path		int		true	"Category ID."
//	@Success		200			{object}	models.CategoryList
//	@Failure		404			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/category/{category} [get]
func get(ctx *gin.Context) {

	var category *models.Category

	categoryList := &models.CategoryList{}

	if err := db.Tx.Where("user_id", ctx.GetUint("id")).First(&category, ctx.Param("category")).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "category not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return

	}
	structure.Assign(category, categoryList)

	ctx.JSON(http.StatusOK, categoryList)

}
