package server

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/v7ktory/test/internal/config"
)

type Server struct {
	logger *slog.Logger

	app *fiber.App
}

func New(cfg *config.Config, logger *slog.Logger, app *fiber.App) *Server {
	return &Server{
		logger: logger,
		app:    app,
	}
}

func (s *Server) MustRun(port string) {
	if err := s.run(port); err != nil {
		panic(err)
	}

}
func (s *Server) run(port string) error {
	s.logger.Info("Http server started", slog.String("addr", port))

	return s.app.Listen(port)
}

func (s *Server) Stop() error {
	s.logger.Info("Stopping http server")
	return s.app.Shutdown()
}
