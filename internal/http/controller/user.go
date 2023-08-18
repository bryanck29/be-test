package controller

import (
	"net/http"

	"github.com/bryanck29/be-test/internal/constant"
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/schema/model"
	"github.com/bryanck29/be-test/internal/schema/request"
	"github.com/bryanck29/be-test/internal/utils"
	_ "github.com/bryanck29/be-test/pkg/model"
	extUtils "github.com/bryanck29/be-test/pkg/utils"

	"github.com/labstack/echo/v4"
)

// userController represents the user controller object
type userController struct {
	userUsecase contract.UserUsecase
}

// newUserController used to intialize user controller
func newUserController(e *echo.Echo, userUsecase contract.UserUsecase) contract.UserController {
	return &userController{
		userUsecase: userUsecase,
	}
}

// PostInsertUser godoc
//
//	@Summary		Insert a new user
//	@Description	Insert a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body			body		request.PostInsertUser	true	"Request body for creating a new user"
//	@Param			Authorization	header		string					true	"Authorization"
//	@Success		201				{object}	model.Response
//	@Failure		400				{object}	model.Response
//	@Failure		401				{object}	model.Response
//	@Failure		500				{object}	model.Response
//	@Router			/user [post]
func (c *userController) PostInsertUser(ctx echo.Context) error {
	req := new(request.PostInsertUser)
	if err := extUtils.ParseParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, constant.ErrParsingRequest, "", nil)
	}

	if err := extUtils.ValidateParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	err := c.userUsecase.InsertUser(ctx, *req)
	if err != nil {
		return utils.ErrorResponse(ctx, err, constant.ErrCreatingUser.Error(), nil)
	}

	return utils.SuccessResponse(ctx, http.StatusCreated, constant.UserCreated, nil)
}

// GetUsers godoc
//
//	@Summary		Get users
//	@Description	Retrieves a list of users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization"
//	@Success		200				{array}		model.User
//	@Failure		401				{object}	model.Response
//	@Failure		500				{object}	model.Response
//	@Router			/user [get]
func (c *userController) GetUsers(ctx echo.Context) error {
	result, err := c.userUsecase.GetUsers(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, err, constant.ErrGettingUser.Error(), nil)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, constant.GetUsers, result)
}

// GetUser godoc
//
//	@Summary		Get user
//	@Description	Retrieves a single user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization"
//	@Param			userId			path		string	true	"User ID should be a valid UUID string"
//	@Success		200				{object}	model.Response
//	@Failure		400				{object}	model.Response
//	@Failure		401				{object}	model.Response
//	@Failure		404				{object}	model.Response
//	@Failure		500				{object}	model.Response
//	@Router			/user/{userId} [get]
func (c *userController) GetUser(ctx echo.Context) error {
	req := new(request.GetUser)
	if err := extUtils.ParseParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, constant.ErrParsingRequest, "", nil)
	}

	if err := extUtils.ValidateParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	user := ctx.Get(constant.ClaimsUser).(model.User)
	if user.Role == constant.USER_ROLE_USER && user.Id != req.UserId {
		err := constant.ErrInvalidAccess
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	result, err := c.userUsecase.GetUser(ctx, *req)
	if err != nil {
		return utils.ErrorResponse(ctx, err, constant.ErrGettingUser.Error(), nil)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, constant.GetUser, result)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Authorization"
//	@Param			userId			path		string	true	"User ID should be a valid UUID string"
//	@Success		200				{object}	model.Response
//	@Failure		400				{object}	model.Response
//	@Failure		401				{object}	model.Response
//	@Failure		500				{object}	model.Response
//	@Router			/user/{userId} [delete]
func (c *userController) DeleteUser(ctx echo.Context) error {
	req := new(request.DeleteUser)
	if err := extUtils.ParseParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, constant.ErrParsingRequest, "", nil)
	}

	if err := extUtils.ValidateParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	err := c.userUsecase.DeleteUser(ctx, *req)
	if err != nil {
		return utils.ErrorResponse(ctx, err, constant.ErrDeletingUser.Error(), nil)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, constant.DeleteUser, nil)
}

// DeleteUser godoc
//
//	@Summary		Updates a user by ID
//	@Description	Updates a user by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//
//	@Param			Authorization	header		string			true	"Authorization"
//	@Param			userId			path		string			true	"User ID should be a valid UUID string"
//
//	@Param			input			body		request.PutUser	true	"Just include the fields you want to update"
//	@Success		200				{object}	model.Response
//	@Failure		400				{object}	model.Response
//	@Failure		401				{object}	model.Response
//	@Failure		500				{object}	model.Response
//	@Router			/user/{userId} [put]
func (c *userController) PutUser(ctx echo.Context) error {
	req := new(request.PutUser)
	if err := extUtils.ParseParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, constant.ErrParsingRequest, "", nil)
	}

	if err := extUtils.ValidateParameter(ctx, req); err != nil {
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	result, err := c.userUsecase.UpdateUser(ctx, *req)
	if err != nil {
		return utils.ErrorResponse(ctx, err, "", nil)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, constant.UpdateUser, result)
}
