package app

import (
	"log"

	"github.com/kuznetsovgg/go-server-practice/internal/app/endpoint"
	"github.com/kuznetsovgg/go-server-practice/internal/app/middlewares"
	"github.com/kuznetsovgg/go-server-practice/internal/app/service"
	"github.com/labstack/echo/v4"
)

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

	a.echo.GET("/", a.e.Check, middlewares.UAcheck)

	return a, nil
}

func (a *App) Run() error {
	err := a.echo.Start(":8080")
	log.Println("Server started!")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
