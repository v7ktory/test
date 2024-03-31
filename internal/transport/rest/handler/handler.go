package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/v7ktory/test/internal/service"
)

type Handler struct {
	log *slog.Logger
	svc *service.Service
}

func New(log *slog.Logger, svc *service.Service) *Handler {
	return &Handler{
		log: log,
		svc: svc,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	r := fiber.New()
	r.Post("/edit/:id", h.EditPost)
	r.Get("/list", h.GetList)
	return r
}
