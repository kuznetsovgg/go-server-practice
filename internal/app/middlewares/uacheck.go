package middlewares

import (
	"log"

	"github.com/labstack/echo/v4"
)

func UAcheck(next echo.HandlerFunc) echo.HandlerFunc {
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
