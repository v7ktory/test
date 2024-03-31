package app

import (
	"log/slog"

	"github.com/v7ktory/test/internal/config"
	"github.com/v7ktory/test/internal/server"
	"github.com/v7ktory/test/internal/service"
	"github.com/v7ktory/test/internal/storage/postgres"
	"github.com/v7ktory/test/internal/transport/rest/handler"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type App struct {
	Server *server.Server
}

func New(logger *slog.Logger, cfg *config.Config) *App {
	p := postgres.New(cfg)
	db := reform.NewDB(p.Pool, postgresql.Dialect, reform.NewPrintfLogger(logger.Info))

	svc := service.New(db, logger)

	h := handler.New(logger, svc)

	r := h.InitRoutes()

	srv := server.New(cfg, logger, r)

	return &App{
		Server: srv,
	}
}
