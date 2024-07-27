package server

import (
	"context"
	"embed"
	"net/http"

	"github.com/bastean/codexgo/internal/app/server/router"
	"github.com/bastean/codexgo/internal/pkg/service/env"
	"github.com/bastean/codexgo/internal/pkg/service/errors"
	"github.com/bastean/codexgo/internal/pkg/service/logger/log"
)

var (
	Server = &struct {
		Gin string
	}{
		Gin: log.Server("Gin"),
	}
)

//go:embed static
var Files embed.FS

var App *http.Server

func Up() error {
	log.Starting(Server.Gin)

	App = &http.Server{
		Addr:    ":" + env.ServerGinPort,
		Handler: router.New(&Files),
	}

	if err := App.ListenAndServe(); err != nil {
		log.CannotBeStarted(Server.Gin)
		return errors.BubbleUp(err, "Up")
	}

	log.Started(Server.Gin)

	return nil
}

func Down(ctx context.Context) error {
	log.Stopping(Server.Gin)

	if err := App.Shutdown(ctx); err != nil {
		log.CannotBeStopped(Server.Gin)
		return errors.BubbleUp(err, "Down")
	}

	log.Stopped(Server.Gin)

	return nil
}
