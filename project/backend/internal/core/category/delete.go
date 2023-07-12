package category

import (
	"net/http"
	"strings"

	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		DELETE
//	@Description	Delete many or all categories.
//	@Tags			category
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			categories	query		[]int	false	"Category ID's."
//	@Success		204			{string}	string	"No Content"
//	@Failure		400			{object}	models.HTTP
//	@Failure		500			{object}	models.HTTP
//	@Router			/category [delete]
func delete(ctx *gin.Context) {

	var result *gorm.DB

	if "" != ctx.Query("categories") {
		result = db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Category{}, strings.Split(ctx.Query("categories"), ","))
	} else {
		result = db.Tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Category{})
	}
	if result.Error != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			result.Error.Error(),
		)
		return
	} else if result.RowsAffected < 1 {
		api.Return(
			ctx,
			-1,
			"no removed",
		)
		return
	}

	ctx.Status(http.StatusNoContent)

}

// Swagger:
//
//	@Summary		DELETE
//	@Description	Delete the category.
//	@Tags			category
//	@Param			TOKEN		header		string	true	"Bearer token."
//	@Param			category	path		int		true	"Category ID."
//	@Success		204			{string}	string	"No Content"
//	@Failure		500			{object}	models.HTTP
//	@Failure		000			{string}	string	"No Removed (-1)"
//	@Router			/category/{category} [delete]
func deleteCategory(ctx *gin.Context) {

	if result := db.Tx.Unscoped().Where("user_id", ctx.GetUint("id")).Delete(&models.Category{}, ctx.Param("category")); result.Error != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			result.Error.Error(),
		)
		return
	} else if result.RowsAffected < 1 {
		api.Return(
			ctx,
			-1,
			"no removed",
		)
		return
	}

	ctx.Status(http.StatusNoContent)

}
