package handler

import "message/src/pkg/logger"

type Handler struct {
	logger logger.Logger
}

func NewHandler(logger logger.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
