package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-architecture/internal/app/endpoint"
	"go-architecture/internal/app/middleware"
	"go-architecture/internal/app/service"
	"log"
	"strconv"
)

const HTTP_PORT = 8080

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.GET("/date", a.e.Status, middleware.DateCheck)

	return a, nil
}

func (a *App) Start() error {
	fmt.Printf("Server running on port %d", HTTP_PORT)

	err := a.echo.Start(":" + strconv.Itoa(HTTP_PORT))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
