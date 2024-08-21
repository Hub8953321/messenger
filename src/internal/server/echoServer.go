package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"messager/src/internal/handler"
	"messager/src/internal/server/midleware"
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

	auth := s.Group("/auth")
	{
		auth.POST("/sign-in", s.handler.SignIn)
		auth.POST("/sign-up", s.handler.SingUp)
		auth.POST("/refresh", s.handler.Refresh)
	}

	chat := s.Group("/chat")
	{
		chat.DELETE("/member", midleware.CheckAuth(s.handler.RemoveMembers, s.handler))
		chat.POST("/member", midleware.CheckAuth(s.handler.AddMembers, s.handler))
		chat.POST("/", midleware.CheckAuth(s.handler.CreateChat, s.handler))
		//TODO chat.DELETE("/", midleware.CheckAuth(s.handler.Re))
	}
}

func (s *EchoServer) Stop() error {
	return s.Shutdown(context.Background())
}
