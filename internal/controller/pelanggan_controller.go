package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"laundrymu-api/internal/service"
	"laundrymu-api/internal/utils"
	"net/http"
)

type PelangganController interface {
	FindByTokoId(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type pelangganController struct {
	PelangganService service.PelangganService
}

func NewPelangganController(db *sql.DB) PelangganController {
	return &pelangganController{
		PelangganService: service.NewPelangganService(db),
	}
}

func (p *pelangganController) FindByTokoId(ctx *fiber.Ctx) error {
	pelanggan := entity.PelangganByTokoId{
		TokoId: ctx.QueryInt("toko_id"),
	}

	errors := utils.Validation(&pelanggan)
	if len(errors) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ResponseError{
			Status:  false,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	result, err := p.PelangganService.FindByTokoId(ctx, pelanggan.TokoId)
	if err != nil {
		return ctx.Status(500).JSON(utils.ResponseError{Status: false, Message: err.Error(), Errors: err})
	}

	return ctx.JSON(utils.ResponseData{Status: true, Message: "Get pelanggan success", Data: result})
}

func (p *pelangganController) Create(ctx *fiber.Ctx) error {
	var pelanggan entity.PelangganRequest

	if err := ctx.BodyParser(&pelanggan); err != nil {
		return ctx.Status(400).JSON(utils.ResponseError{Status: false, Message: "JSON format failed"})
	}

	errors := utils.Validation(&pelanggan)
	if len(errors) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ResponseError{
			Status:  false,
			Message: "Validation error",
			Errors:  errors,
		})
	}

	err := p.PelangganService.Create(ctx, &pelanggan)
	if err != nil {
		return ctx.Status(400).JSON(utils.ResponseError{
			Status:  false,
			Message: "Create pelanggan failed",
			Errors:  err,
		})
	}

	return ctx.JSON(utils.ResponseSuccess{Status: true, Message: "Create pelanggan success"})
}
