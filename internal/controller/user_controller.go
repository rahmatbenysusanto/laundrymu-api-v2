package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"laundrymu-api/internal/service"
	"laundrymu-api/internal/utils"
	"net/http"
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type userController struct {
	UserService service.UserService
}

func NewUserController(db *sql.DB) UserController {
	return &userController{
		UserService: service.NewUserService(db),
	}
}

func (u *userController) Login(ctx *fiber.Ctx) error {
	var user = entity.UserLogin{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(utils.ResponseError{Status: false, Message: "JSON format failed"})
	}

	errors := utils.Validation(&user)
	if len(errors) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ResponseError{
			Status:  false,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	result, err := u.UserService.Login(ctx, &user)
	if err != nil {
		return ctx.Status(400).JSON(utils.ResponseError{
			Status:  false,
			Message: "Login Failed",
			Errors:  err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ResponseData{
		Status:  true,
		Message: "Login successfully",
		Data:    result,
	})
}

func (u *userController) Register(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
