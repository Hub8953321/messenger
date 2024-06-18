package main

import (
	"message/src/internal/handler"
	"message/src/internal/server"
	"message/src/pkg/logger"
)

func main() {
	logger := logger.NewTempLogger()
	handler := handler.NewHandler(logger)
	server := server.NewEchoServer()
	server.InitRoutes(handler)
	server.Run(":8080")
}
