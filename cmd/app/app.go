package app

import (
	"DataTask/internal/config"
	"DataTask/internal/di"
	"fmt"
)

func Run(cfg *config.Config) error {
	app, err := di.InitializeApp(cfg)
	if err != nil {
		return err
	}

	fmt.Println(app.Config.RabbitMQ.User)

	return nil
}
