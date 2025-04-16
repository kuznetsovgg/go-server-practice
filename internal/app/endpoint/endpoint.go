package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	s Service
}

type Service interface {
	CountDays() int64
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Check(c echo.Context) error {
	days := e.s.CountDays()

	msg := fmt.Sprintf("Количество дней: %d", days)

	err := c.String(http.StatusOK, msg)
	if err != nil {
		return err
	}

	return nil
}
