package controller

import (
	"net/http"

	"github.com/bryanck29/be-test/internal/constant"
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/schema/request"
	"github.com/bryanck29/be-test/internal/utils"
	_ "github.com/bryanck29/be-test/pkg/model"
	extUtils "github.com/bryanck29/be-test/pkg/utils"

	"github.com/labstack/echo/v4"
)

// authController represents the auth controller object
type authController struct {
	authUsecase contract.AuthUsecase
}

// newAuthController used to intialize auth controller
func newAuthController(e *echo.Echo, authUsecase contract.AuthUsecase) contract.AuthController {
	return &authController{
		authUsecase: authUsecase,
	}
}

// PostLogin godoc
//
// @Summary		User login
// @Description	Authenticate user and return token
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Param			request	body		request.PostLogin	true	"User login parameters"
// @Success		200		{object}	response.PostLogin
// @Failure		400		{object}	model.Response
// @Failure		500		{object}	model.Response
// @Router			/auth/login [post]
func (c *authController) PostLogin(ctx echo.Context) error {
	req := new(request.PostLogin)
	if err := extUtils.ParseParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, constant.ErrParsingRequest, "", nil)
	}

	if err := extUtils.ValidateParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	result, err := c.authUsecase.Login(ctx, *req)
	if err != nil {
		return utils.ErrorResponse(ctx, err, constant.ErrLogin.Error(), nil)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, constant.LoginSucceed, result)
}

// PostRefreshLogin godoc
//
// @Summary		User login
// @Description	Rerfresh user session token
// @Tags			Authentication
// @Accept			json
// @Produce		json
// @Param			Authorization	header		string						true	"Authorization"
// @Param			request			body		request.PostRefreshLogin	true	"User session token"
// @Success		200				{object}	response.PostRefreshLogin
// @Failure		400				{object}	model.Response
// @Failure		500				{object}	model.Response
// @Router			/auth/refresh [post]
func (c *authController) PostRefreshLogin(ctx echo.Context) error {
	req := new(request.PostRefreshLogin)
	if err := extUtils.ParseParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, constant.ErrParsingRequest, "", nil)
	}

	if err := extUtils.ValidateParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	result, err := c.authUsecase.RefreshLogin(ctx, *req)
	if err != nil {
		return utils.ErrorResponse(ctx, err, constant.ErrRefreshLogin.Error(), nil)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, constant.RefreshSucceed, result)
}
