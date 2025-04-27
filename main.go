package main

import (
	"DataTask/cmd/app"
	"DataTask/internal/config"
)

func main() {
	cfg, err := config.NewConfig(
		"./infra/config/.env",
		"./infra/config",
		"config",
	)
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
