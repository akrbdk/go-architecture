package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func DateCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		curDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
		log.Printf("Current date: %s", curDate)

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
