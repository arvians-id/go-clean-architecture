package controller

import (
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/request"
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/response"
	"github.com/arvians-id/go-clean-architecture/internal/service"
	"github.com/arvians-id/go-clean-architecture/util"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserServiceContract
}

func NewUserController(userService service.UserServiceContract) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	users, err := controller.UserService.FindAll(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "OK", users)
}

func (controller *UserController) FindByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := controller.UserService.FindByID(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "OK", user)
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	var userRequest request.CreateUserRequest
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	errValidation := util.ValidateStruct(userRequest)
	if errValidation != nil {
		return response.ReturnErrorValidation(ctx, fiber.StatusBadRequest, errValidation)
	}

	user, err := controller.UserService.Create(ctx.Context(), &userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusCreated, "created", user)
}

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	var userRequest request.UpdateUserRequest
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	errValidation := util.ValidateStruct(userRequest)
	if errValidation != nil {
		return response.ReturnErrorValidation(ctx, fiber.StatusBadRequest, errValidation)
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userRequest.ID = int64(id)
	user, err := controller.UserService.Update(ctx.Context(), &userRequest)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "updated", user)
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = controller.UserService.Delete(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "deleted", nil)
}
