package auth

import (
	"errors"
	"net/http"
	"time"

	"backend/internal/consts"
	"backend/internal/infra/db"
	"backend/internal/models"
	"backend/pkg/utils/api"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt/v2"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		LOGIN
//	@Description	Log-in and get a authentication token (JWT).
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			JSON	body		models.Auth	true	"Json request."
//	@Success		200		{object}	models.Token
//	@Failure		401		{object}	models.HTTP
//	@Failure		404		{object}	models.HTTP
//	@Failure		422		{object}	models.HTTP
//	@Failure		500		{object}	models.HTTP
//	@Router			/auth [post]
func login(ctx *gin.Context) {

	var (
		user    *models.User
		request *models.Auth
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		api.LogReturn(
			ctx,
			http.StatusUnprocessableEntity,
			"malformed JSON",
			err.Error(),
		)
		return
	}

	if err := db.Tx.Where("username", &request.Username).Or("email", &request.Email).First(&user).Error; err != nil {

		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			message = "user not found"
		}

		api.LogReturn(
			ctx,
			code,
			message,
			err.Error(),
		)
		return
	}

	if !user.IsValidPassword(*request.Password) {
		api.Return(ctx, http.StatusUnauthorized, "invalid password")
		return
	}

	token, err := jwt.NewStdToken(
		*user.ID,
		time.Now().Add(time.Duration(consts.JWTTIME)*time.Minute),
		consts.JWTSECRETKEY,
	)

	if err != nil {
		api.LogReturn(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			err.Error(),
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		&models.Token{
			ID:    *user.ID,
			Token: token,
		},
	)

}
