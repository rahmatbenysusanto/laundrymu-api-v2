package route

import (
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/controller"
	"laundrymu-api/internal/middleware"
	"laundrymu-api/pkg/database"
)

var db = database.MysqlConn()

func PublicAPI(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		//go kafka.Producer()
		return c.SendString("Hello, World!")
	})

	pelangganController := controller.NewPelangganController(db)
	app.Get("/pelanggan", pelangganController.FindByTokoId)
	app.Post("/pelanggan", pelangganController.Create)

	chatController := controller.NewChatController(db)
	app.Get("/chat", chatController.FindByTokoId)
	app.Post("/chat", chatController.Create)

	userController := controller.NewUserController(db)
	app.Post("/login", userController.Login)
}

func PrivateAPI(app *fiber.App) {
	app.Use(middleware.AuthJWT)

	app.Get("/tes", func(ctx *fiber.Ctx) error {
		return ctx.SendString("TES JWT")
	})
}
