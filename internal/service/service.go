package service

import (
	"log/slog"

	"gopkg.in/reform.v1"
)

type Service struct {
	logger *slog.Logger
	db     *reform.DB
}

func New(db *reform.DB, logger *slog.Logger) *Service {
	return &Service{db: db, logger: logger}
}
