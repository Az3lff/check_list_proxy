package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/Az3lff/check_list_proxy/internal/config"
	"github.com/Az3lff/check_list_proxy/internal/delivery/grpc"
	"github.com/Az3lff/check_list_proxy/internal/delivery/http"
	"github.com/Az3lff/check_list_proxy/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("cannot load environment: %s", err.Error())
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err.Error())
	}

	cli, err := grpc.New(cfg.GRPCClient)
	if err != nil {
		logrus.Fatalf("cannot to create grpc client: %s", err.Error())
	}

	svs := service.NewService(cli)

	app := fiber.New()
	http.SetupRoutes(app, cfg.HTTPServer, svs)
	logrus.Fatal(app.Listen(fmt.Sprintf(":" + os.Getenv("PORT"))))
}
