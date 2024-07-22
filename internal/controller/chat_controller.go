package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"laundrymu-api/internal/service"
	"laundrymu-api/internal/utils"
	"net/http"
)

type ChatController interface {
	FindByTokoId(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type chatController struct {
	ChatService service.ChatService
}

func NewChatController(db *sql.DB) ChatController {
	return &chatController{
		ChatService: service.NewChatService(db),
	}
}

func (c *chatController) FindByTokoId(ctx *fiber.Ctx) error {
	chat := entity.PelangganByTokoId{
		TokoId: ctx.QueryInt("toko_id"),
	}

	errors := utils.Validation(&chat)
	if len(errors) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ResponseError{
			Status:  false,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	result, err := c.ChatService.FindByTokoId(ctx, chat.TokoId)

	if err != nil {
		return ctx.Status(500).JSON(utils.ResponseError{Status: false, Message: err.Error(), Errors: err})
	}

	return ctx.JSON(utils.ResponseData{Status: true, Message: "Get chat success", Data: result})
}

func (c *chatController) Create(ctx *fiber.Ctx) error {
	var chat = entity.ChatCreate{}
	if err := ctx.BodyParser(&chat); err != nil {
		return ctx.Status(400).JSON(utils.ResponseError{Status: false, Message: "JSON format failed"})
	}

	errors := utils.Validation(&chat)
	if len(errors) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ResponseError{
			Status:  false,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	err := c.ChatService.Create(ctx, &chat)
	if err != nil {
		return ctx.Status(400).JSON(utils.ResponseError{
			Status:  false,
			Message: "Create chat failed",
			Errors:  err,
		})
	}

	return ctx.JSON(utils.ResponseSuccess{Status: true, Message: "Create chat success"})
}
