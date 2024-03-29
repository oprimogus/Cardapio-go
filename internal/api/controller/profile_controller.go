package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	validatorutils "github.com/oprimogus/cardapiogo/internal/api/validator"
	"github.com/oprimogus/cardapiogo/internal/domain/profile"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

type ProfileController struct {
	service   *profile.Service
	validator *validatorutils.Validator
}

func NewProfileController(repository profile.Repository, userRepository user.Repository, validator *validatorutils.Validator) *ProfileController {
	return &ProfileController{service: profile.NewService(repository, userRepository), validator: validator}
}

// CreateProfileHandler godoc
// @Summary Cria um perfil e atribui a um usuário
// @Description Cria um perfil e atribui a um usuário
// @Tags Profile
// @Accept  json
// @Produce  json
// @Param   request body profile.CreateProfileParams true "CreateProfileParams"
// @Success 201
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /v1/profile [post]
func (c *ProfileController) CreateProfileHandler(ctx *gin.Context) {
	var createProfileParams profile.CreateProfileParams
	err := ctx.BindJSON(&createProfileParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.BadRequest(err.Error()))
		return
	}

	validationErr := c.validator.Validate(createProfileParams)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, validationErr)
	}

	userID := ctx.GetString("userID")

	err = c.service.CreateProfile(ctx, userID, createProfileParams)
	validateErrorResponse(ctx, err)
	ctx.Status(http.StatusOK)
}

// GetProfileByUserIDHandler godoc
// @Summary Retorna o perfil do usuário
// @Description Retorna o perfil do usuário
// @Tags Profile
// @Accept  json
// @Produce  json
// @Success 200  {object} profile.Profile
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /v1/profile [get]
func (c *ProfileController) GetProfileByUserIDHandler(ctx *gin.Context) {
	userID := ctx.GetString("userID")

	profile, err := c.service.GetProfileByUserID(ctx, userID)
	validateErrorResponse(ctx, err)
	ctx.JSON(http.StatusOK, profile)
}

// UpdateProfileHandler godoc
// @Summary Atualiza os dados do perfil do usuário
// @Description Atualiza os dados do perfil do usuário
// @Tags Profile
// @Accept  json
// @Produce  json
// @Param   request body profile.UpdateProfileParams true "UpdateProfileParams"
// @Success 200
// @Failure 400  {object} errors.ErrorResponse
// @Failure 500  {object} errors.ErrorResponse
// @Failure 502  {object} errors.ErrorResponse
// @Router /v1/profile [put]
func (c *ProfileController) UpdateProfileHandler(ctx *gin.Context) {

	var updateProfileParams profile.UpdateProfileParams
	err := ctx.BindJSON(&updateProfileParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.BadRequest(err.Error()))
		return
	}

	validationErr := c.validator.Validate(updateProfileParams)
	if validationErr != nil {
		ctx.JSON(validationErr.Status, validationErr)
		return
	}

	userID := ctx.GetString("userID")

	err = c.service.UpdateProfile(ctx, userID, updateProfileParams)
	validateErrorResponse(ctx, err)

	ctx.Status(http.StatusOK)
}
