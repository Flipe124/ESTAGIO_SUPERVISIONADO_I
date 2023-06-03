package order

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
//	@Description	List all Orders.
//	@Tags			order
//	@Produce		json
//	@Param			Token	header		string	true	"Bearer token."
//	@Success		200		{array}		models.OrderList
//	@Failure		500		{object}	models.HTTP
//	@Router			/order [get]
func list(ctx *gin.Context) {

	var orders []*models.Order

	if err := db.Instance.
		Preload("User").Select("id", "name").
		Preload("Client").Select("id", "name").
		Preload("Service").
		Find(&orders).Error; err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ordersList := make([]*models.OrderList, len(orders))
	for index, order := range orders {
		ordersList[index] = &models.OrderList{}
		structure.Assign(order, ordersList[index], "User", "Client", "Service")
		structure.Assign(order.User, ordersList[index].User)
		structure.Assign(order.Client, ordersList[index].Client)
		structure.Assign(order.Service, ordersList[index].Service)
	}

	ctx.JSON(http.StatusOK, ordersList)

}
