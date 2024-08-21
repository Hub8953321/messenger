package handler

import (
	"github.com/Eugene-Usachev/fst"
	"messager/src/internal/service"
	"messager/src/pkg/logger"
)

type Handler struct {
	logger.Logger
	*service.Service
	AccessConverter  *fst.EncodedConverter
	RefreshConverter *fst.EncodedConverter
}

func NewHandler(
	logger logger.Logger,
	service *service.Service,
	accessConverter,
	refreshConverter *fst.EncodedConverter,
) *Handler {
	return &Handler{
		logger,
		service,
		accessConverter,
		refreshConverter,
	}
}
