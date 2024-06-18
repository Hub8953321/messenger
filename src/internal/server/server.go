package server

import "message/src/internal/handler"

type Server interface {
	Run(string) error
	InitRoutes(*handler.Handler)
	Stop() error
}
