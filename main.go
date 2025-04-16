package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", Handler)
	log.Println("Server started!")
	e.Logger.Fatal(e.Start(":8080"))
}

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
