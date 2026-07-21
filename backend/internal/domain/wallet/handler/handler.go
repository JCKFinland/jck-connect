package handler

import (
	walletservice "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/service"
)

type Handler struct {
	service walletservice.Service
}

func New(
	service walletservice.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}
