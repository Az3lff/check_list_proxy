package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Az3lff/check_list_proxy/internal/config"
	"github.com/Az3lff/check_list_proxy/internal/service"
)

func SetupRoutes(app *fiber.App, cfg config.HTTPServer, svs *service.Service) {
	h := NewHandler(cfg, svs)

	app.Get("api/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("ok")
	})

	app.Post("/api/create", h.CreateTask)
	app.Get("/api/list", h.GetList)
	app.Delete("/api/delete", h.DeleteTask)
	app.Put("/api/done", h.DoneTask)
}
