package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wmattei/gotodo/pkg/metadata"
	"github.com/wmattei/gotodo/pkg/services"
)

func getApiMetadata(ctx *fiber.Ctx) error {
	return ctx.JSON(metadata.ApiMetadata{
		App:     "GO TODO",
		Version: "0.0.1",
	})
}

func Handler(service *services.TaskService) *fiber.App {
	app := fiber.New()

	app.Get("/", getApiMetadata)
	app.Post("/", PostTask(service))

	return app
}
