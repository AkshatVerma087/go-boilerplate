package handler

import (
	"github.com/AkshatVerma087/go-boilerplate/backend/internal/server"
	"github.com/AkshatVerma087/go-boilerplate/backend/internal/service"
)

type Handlers struct {
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{}
}


