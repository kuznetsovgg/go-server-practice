package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(MW)
	e.GET("/", Handler)
	log.Println("Server started!")

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(c echo.Context) error {
	d := time.Date(2026, 01, 01, 0, 0, 0, 0, time.Local)
	dur := time.Until(d)
	msg := fmt.Sprintf("Количество дней: %d", int64(dur.Hours())/24)

	err := c.String(http.StatusOK, msg)
	if err != nil {
		return err
	}

	return nil
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		val := c.Request().Header.Get("User-Agent")
		log.Println(val)

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
