package main

import (
	"chatmoon/internal/config"
	"fmt"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewGORM(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	webHost := viperConfig.GetString("web.host")
	webPort := viperConfig.GetString("web.port")
	err := app.Listen(fmt.Sprintf("%s:%s", webHost, webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
