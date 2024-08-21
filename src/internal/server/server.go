package server

import "messager/src/internal/handler"

type Server interface {
	Run(string) error
	InitRoutes(*handler.Handler)
	Stop() error
}
