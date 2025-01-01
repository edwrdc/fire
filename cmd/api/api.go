package main

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type application struct {
	config config
	logger *zap.SugaredLogger
}

type config struct {
	addr string
	env  string
}

func (app *application) run() error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.routes(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
