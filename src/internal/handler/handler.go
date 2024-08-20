package handler

import (
	"messager/src/internal/service"
	"messager/src/pkg/logger"
)

type Handler struct {
	logger.Logger
	*service.Service
}

func NewHandler(logger logger.Logger, service *service.Service) *Handler {
	return &Handler{
		logger,
		service,
	}
}
