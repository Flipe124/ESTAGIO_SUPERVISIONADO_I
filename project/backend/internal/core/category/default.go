package category

import (
	"errors"
	"net/http"

	"backend/internal/core/user"
	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/helpers/structure"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		DEFAULT
//	@Description	List the default categories.
//	@Tags			category
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Success		200		{array}		models.CategoryList
//	@Success		204		{string}	string	"No Content"
//	@Failure		500		{object}	models.HTTP
//	@Router			/category/default [get]
func defaultList(ctx *gin.Context) {

	var categories []*models.Category

	if err := db.Tx.Where("user_id", user.SystemUserID).Find(&categories).Error; err != nil {
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

// Swagger:
//
//	@Summary		DEFAULT
//	@Description	List a single default category from ID.
//	@Tags			category
//	@Produce		json
//	@Param			TOKEN	header		string	true	"Bearer token."
//	@Param			default	path		int		true	"Default category ID."
//	@Success		200		{array}		models.CategoryList
//	@Failure		404		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/category/default/{default} [get]
func defaultGet(ctx *gin.Context) {

	var category *models.Category

	categoriesList := &models.CategoryList{}

	if err := db.Tx.Where("user_id", user.SystemUserID).Find(&category, ctx.Param("default")).Error; err != nil {

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
	structure.Assign(category, categoriesList)

	ctx.JSON(http.StatusOK, categoriesList)

}
