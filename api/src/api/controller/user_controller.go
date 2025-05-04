package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"golang-core/api/src/client/exception"
	"golang-core/api/src/client/request"
	"golang-core/api/src/client/response"
	echoutils "golang-core/api/src/common/echo_utils"
	"golang-core/api/src/manager"
)

type UserController struct {
	userManager manager.UserManager
}

func NewUserController(userManager manager.UserManager) *UserController {
	return &UserController{
		userManager: userManager,
	}
}

func (c *UserController) Insert(ec echo.Context) error {
	ctx := ec.Request().Context()
	var req request.UserInsertRequest
	err := echoutils.BindAndValidate(ec, &req)
	if err != nil {
		return exception.NewBadRequestError(err)
	}
	res, err := c.userManager.Insert(ctx, req)
	if err != nil {
		return exception.NewBadRequestError(err)
	}
	return ec.JSON(http.StatusOK, response.ToSuccessResponse(res))
}

func (c *UserController) Update(ec echo.Context) error {
	ctx := ec.Request().Context()
	var req request.UserUpdateRequest
	err := echoutils.BindAndValidate(ec, &req)
	if err != nil {
		return exception.NewBadRequestError(err)
	}
	res, err := c.userManager.Update(ctx, req)
	if err != nil {
		return exception.NewBadRequestError(err)
	}
	if res == nil {
		return exception.NewBadRequestError(errors.New("No record found."))
	}
	return ec.JSON(http.StatusOK, response.ToSuccessResponse(res))
}

func (c *UserController) Delete(ec echo.Context) error {
	ctx := ec.Request().Context()
	id := ec.Param("id")
	res, err := c.userManager.FindById(ctx, id)
	if err != nil {
		return exception.NewBadRequestError(err)
	}
	err = c.userManager.Delete(ctx, *res.Id)
	return ec.JSON(http.StatusOK, response.ToSuccessResponse("Deleted"))
}

func (c *UserController) FindById(ec echo.Context) error {
	ctx := ec.Request().Context()
	id := ec.Param("id")
	res, err := c.userManager.FindById(ctx, id)
	if err != nil {
		return exception.NewBadRequestError(err)
	}
	return ec.JSON(http.StatusOK, response.ToSuccessResponse(res))
}

func (c *UserController) ListByIds(ec echo.Context) error {
	ctx := ec.Request().Context()
	var req request.UserListAllRequest
	if err := echoutils.BindAndValidate(ec, &req); err != nil {
		return exception.NewBadRequestError(err)
	}
	res, err := c.userManager.ListByIds(ctx, *req.Limit, *req.Offset)
	if err != nil {
		return exception.NewBadRequestError(err)
	}
	return ec.JSON(http.StatusOK, response.ToSuccessResponse(res))
}
