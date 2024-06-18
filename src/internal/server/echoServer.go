package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"message/src/internal/handler"
)

type EchoServer struct {
	*echo.Echo
	handler *handler.Handler
}

var _ Server = (*EchoServer)(nil)

func NewEchoServer() *EchoServer {
	return &EchoServer{
		echo.New(),
		nil,
	}
}

func (s *EchoServer) Run(address string) error {
	return s.Start(address)
}

func (s *EchoServer) InitRoutes(handler *handler.Handler) {
	s.handler = handler
	s.POST("/auth/sign-in", s.handler.SingIn)
}

func (s *EchoServer) Stop() error {
	return s.Shutdown(context.Background())
}
