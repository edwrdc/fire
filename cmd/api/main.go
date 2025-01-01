package main

import (
	"github.com/edwrdc/fire/internal/env"
	"go.uber.org/zap"
)

const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.Get("ADDR", ":8080"),
		env:  env.Get("ENV", "development"),
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	app := &application{
		config: cfg,
		logger: logger,
	}

	app.logger.Infow("Server Started", "env", app.config.env, "addr", app.config.addr)

	logger.Fatal(app.run())
}
