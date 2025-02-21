package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/Az3lff/check_list_proxy/internal/config"
	"github.com/Az3lff/check_list_proxy/internal/delivery/grpc"
	"github.com/Az3lff/check_list_proxy/internal/delivery/http"
	"github.com/Az3lff/check_list_proxy/internal/service"
)

func Run(cfg *config.Config) error {
	app := fiber.New()

	cli, err := grpc.New(cfg.GRPCClient)
	if err != nil {
		return fmt.Errorf("cannot to create grpc client: %w", err)
	}

	svs := service.NewService(cli)
	http.SetupRoutes(app, cfg.HTTPServer, svs)

	return app.Listen(cfg.HTTPServer.Address)
}
