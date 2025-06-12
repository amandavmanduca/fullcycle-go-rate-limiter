package handlers

import (
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/containers/service"
	"github.com/amandavmanduca/fullcycle-go-rate-limiter/src/handlers/middlewares"
	"github.com/labstack/echo/v4"
)

type handler struct {
	services service.ServiceContainer
}

func NewHandler(services service.ServiceContainer) *handler {
	return &handler{
		services: services,
	}
}

func Start(services service.ServiceContainer) {
	e := echo.New()

	middlewares := middlewares.NewMiddleware(services)

	e.Use(middlewares.RateLimitMiddleware)

	h := NewHandler(services)
	e.GET("/", h.Hello)
	e.Start(":8080")
}
