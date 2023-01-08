package controller

import (
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/request"
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/response"
	"github.com/arvians-id/go-clean-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	users, err := controller.UserService.FindAll(ctx)
	if err != nil {
		response.ReturnErrorInternalServerError(ctx, err, nil)
		return
	}

	response.ReturnSuccessOK(ctx, "OK", users)
}

func (controller *UserController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(ctx, err, nil)
		return
	}

	user, err := controller.UserService.FindByID(ctx, id)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			response.ReturnErrorNotFound(ctx, err, nil)
			return
		}
		response.ReturnErrorInternalServerError(ctx, err, nil)
		return
	}

	response.ReturnSuccessOK(ctx, "OK", user)
}

func (controller *UserController) Create(ctx *gin.Context) {
	var userRequest request.CreateUserRequest
	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		response.ReturnErrorBadRequest(ctx, err, nil)
		return
	}

	user, err := controller.UserService.Create(ctx, &userRequest)
	if err != nil {
		response.ReturnErrorInternalServerError(ctx, err, nil)
		return
	}

	response.ReturnSuccessOK(ctx, "OK", user)
}

func (controller *UserController) Update(ctx *gin.Context) {
	var userRequest request.UpdateUserRequest
	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		response.ReturnErrorBadRequest(ctx, err, nil)
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(ctx, err, nil)
		return
	}

	userRequest.ID = id
	user, err := controller.UserService.Update(ctx, &userRequest)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			response.ReturnErrorNotFound(ctx, err, nil)
			return
		}
		response.ReturnErrorInternalServerError(ctx, err, nil)
		return
	}

	response.ReturnSuccessOK(ctx, "OK", user)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.ReturnErrorBadRequest(ctx, err, nil)
		return
	}

	err = controller.UserService.Delete(ctx, id)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			response.ReturnErrorNotFound(ctx, err, nil)
			return
		}
		response.ReturnErrorInternalServerError(ctx, err, nil)
		return
	}

	response.ReturnSuccessOK(ctx, "OK", nil)
}
