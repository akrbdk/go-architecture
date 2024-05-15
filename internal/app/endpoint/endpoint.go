package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() (int, int64)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) Status(c echo.Context) error {
	nextYear, d := e.s.DaysLeft()
	msg := fmt.Sprintf("Days until New %d Year: %d", nextYear, d)
	err := c.String(http.StatusOK, msg)
	if err != nil {
		return err
	}

	return nil
}
