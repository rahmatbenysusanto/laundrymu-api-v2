package middleware

import (
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/service"
	"laundrymu-api/internal/utils"
)

func AuthJWT(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")
	if auth == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError{
			Status:  false,
			Message: "Unauthorized",
			Errors:  "Token not provided",
		})
	}

	err := service.ValidateToken(auth)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError{
			Status:  false,
			Message: "Invalid token",
			Errors:  err,
		})
	}

	err = ctx.Next()
	if err != nil {
		return err
	}
	return nil
}
