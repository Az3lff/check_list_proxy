package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Az3lff/check_list_proxy/internal/app"
	"github.com/Az3lff/check_list_proxy/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err.Error())
	}

	// TODO: add custom logger

	// TODO: add observability

	if err := app.Run(cfg); err != nil {
		logrus.Fatal(err.Error())
	}
}
