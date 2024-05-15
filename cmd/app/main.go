package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server running on port 8080")

	e := echo.New()
	e.GET("/date", Handler, DateMiddleware)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(c echo.Context) error {
	nextYear := time.Now().Year() + 1
	durToNewYear := time.Until(time.Date(nextYear, time.January, 1, 0, 0, 0, 0, time.UTC))
	msg := fmt.Sprintf("Days until New %d Year: %d", nextYear, int64(durToNewYear.Hours())/24)
	err := c.String(http.StatusOK, msg)
	if err != nil {
		return err
	}

	return nil
}

func DateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
