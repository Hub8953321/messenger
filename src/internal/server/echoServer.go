package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"messager/src/internal/handler"
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
	//TODO
	s.POST("/auth/sign-in", s.handler.SignIn)
}

func (s *EchoServer) Stop() error {
	return s.Shutdown(context.Background())
}
